package service

import (
	"encoding/json"
	"fmt"
	"git.ont.io/ontid/otf/config"
	"git.ont.io/ontid/otf/did"
	"git.ont.io/ontid/otf/message"
	"git.ont.io/ontid/otf/middleware"
	"git.ont.io/ontid/otf/store"
	"github.com/fatih/structs"
	"github.com/google/uuid"
	"github.com/ontio/ontology-crypto/signature"
	sdk "github.com/ontio/ontology-go-sdk"
)

const (
	Version        = "1.0"
	InvitationSpec = "spec/connections/" + Version + "/invitation"
	ConnectionRequest = "spec/connections/" + Version + "/request"
	InvitationKey  = "Invitation"
	ConnectionKey  = "Connection"
)

type Syscontroller struct {
	account *sdk.Account
	did     did.Did
	cfg     *config.Cfg
	store   store.Store
}

func NewSyscontroller(acct *sdk.Account, cfg *config.Cfg, db store.Store) Syscontroller {
	did := did.NewOntDID(cfg, acct)
	s := Syscontroller{
		account: acct,
		did:     did,
		cfg:     cfg,
		store:   db,
	}
	s.Initiate(nil)
	return s
}

func (s Syscontroller) Name() string {
	return "syscontroller"
}

func (s Syscontroller) Initiate(param ParameterInf) error {
	fmt.Printf("%s Initiate\n", s.Name())
	//todo add logic
	return nil
}

func (s Syscontroller) Process(msg message.Message) (ControllerResp, error) {
	fmt.Printf("%s Process:%v\n", s.Name(), msg)
	//todo add logic
	switch msg.MessageType {
	//for system
	case message.InvitationType:
		middleware.Log.Infof("resolve invitation")
		if msg.Content == nil {
			return nil, fmt.Errorf("message content is nil")
		}
		//todo verify request
		invitation, ok := msg.Content.(*message.Invitation)
		if !ok {
			return nil, fmt.Errorf("message format is not correct")
		}
		//set uuid to invitation id
		invitation.Id = uuid.New().String()

		//store the invitation
		err := s.SaveInvitation(*invitation)
		if err != nil {
			return nil, err
		}

		return nil, nil
	case message.SendConnectionRequestType:
		//send connection req for agent
		middleware.Log.Infof("resolve send connection request")
		//todo verify request
		if msg.Content == nil {
			return nil, fmt.Errorf("message content is nil")
		}
		cr := msg.Content.(*message.ConnectionRequest)
		cr.Id = uuid.New().String()
		err := s.SaveConnectionRequest(*cr, ConnectionRequestSent)
		if err != nil {
			return nil, err
		}

		//send the connection req to target service endpoint
		//go handleOutbound(cr)

		//no need to pass incoming param
		return nil, nil

	case message.ConnectionRequestType:
		middleware.Log.Infof("resolve connection request")
		if msg.Content == nil {
			return nil, fmt.Errorf("message content is nil")
		}
		req := msg.Content.(*message.ConnectionRequest)
		//ivid := req.Thread.ID
		ivrc, err := s.GetInvitation(req.Thread.ID)
		if err != nil {
			return nil, err
		}
		//update invitation to used state
		err = s.UpdateInvitation(ivrc.Invitation.Id, InvitationUsed)
		if err != nil {
			return nil, err
		}
		//update connection to request received state
		err = s.SaveConnectionRequest(*req, ConnectionRequestReceived)
		//send response outbound
		res := new(message.ConnectResponse)
		res.Id = uuid.New().String()
		res.Thread = message.Thread{
			ID: req.Id,
		}
		//todo define the response type
		res.Type = ""
		res.Connection = req.Connection

		//todo
		//go outbound(res)
		return nil, nil

	case message.ConnectionResponseType:
		middleware.Log.Infof("resolve connection response")
		//req := msg.Content.(message.ConnectResponse)

		//1. update connection request to receive response state

		//2. create and save a connection object

		//3. send ACK back

	case message.ConnectionACKType:
		middleware.Log.Infof("resolve ConnectionACK")
		//req := msg.Content.(message.ConnectResponse)
		//1. update connection request to receive ack state

		//2. create and save a connection object

	//for custom
	case message.ProposalCredentialType:
	case message.OfferCredentialType:
	case message.RequestCredentialType:
	case message.IssueCredentialType:
	case message.CredentialACKType:

	case message.RequestPresentationType:
	case message.PresentationType:
	case message.PresentationACKType:

	default:

	}

	resp := ServiceResp{}
	return resp, nil
}
func (s Syscontroller) Shutdown() error {
	fmt.Printf("%s shutdown\n", s.Name())
	return nil
}

func (s Syscontroller) sign(data []byte) ([]byte, error) {
	sig, err := signature.Sign(signature.SHA256withECDSA, s.account.PrivateKey, data, nil)
	if err != nil {
		return nil, err
	}
	return signature.Serialize(sig)
}

func (s Syscontroller) toMap(v interface{}) (map[string]interface{}, error) {
	return structs.Map(v), nil
}

//

func (s Syscontroller) SaveInvitation(iv message.Invitation) error {

	key := fmt.Sprintf("%s_%s", InvitationKey, iv.Id)
	b, err := s.store.Has([]byte(key))
	if err != nil {
		return err
	}
	if b {
		return fmt.Errorf("invitation with id:%s existed", iv.Id)
	}

	rec := InvitationRec{
		Invitation: iv,
		State:      InvitationInit,
	}

	bs, err := json.Marshal(rec)
	if err != nil {
		return err
	}

	return s.store.Put([]byte(key), bs)
}

func (s Syscontroller) GetInvitation(id string) (*InvitationRec, error) {
	key := []byte(fmt.Sprintf("%s_%s", InvitationKey, id))
	data, err := s.store.Get(key)
	if err != nil {
		return nil, err
	}

	rec := new(InvitationRec)

	err = json.Unmarshal(data, rec)
	if err != nil {
		return nil, err
	}
	return rec, nil
}

func (s Syscontroller) UpdateInvitation(id string, state ConnectionState) error {
	key := []byte(fmt.Sprintf("%s_%s", InvitationKey, id))
	data, err := s.store.Get(key)
	if err != nil {
		return err
	}
	rec := new(InvitationRec)
	err = json.Unmarshal(data, rec)
	if err != nil {
		return err
	}
	//fixme introduce some FSM
	if rec.State >= state {
		return fmt.Errorf("error state with id:%s", id)
	}
	rec.State = state
	bts, err := json.Marshal(rec)
	if err != nil {
		return err
	}
	return s.store.Put(key, bts)
}

func (s Syscontroller) SaveConnectionRequest(cr message.ConnectionRequest, state ConnectionState) error {
	key := []byte(fmt.Sprintf("%s_%s", ConnectionKey, cr.Id))
	b, err := s.store.Has(key)
	if err != nil {
		return err
	}
	if b {
		return fmt.Errorf("connection request with id:%s existed", cr.Id)
	}
	rec := ConnectionRequestRec{
		ConnReq: cr,
		State:   state,
	}

	bs, err := json.Marshal(rec)
	if err != nil {
		return err
	}

	return s.store.Put(key, bs)
}

func (s Syscontroller) UpdateConnectionRequest(id string, state ConnectionState) error {
	key := []byte(fmt.Sprintf("%s_%s", ConnectionKey, id))
	data, err := s.store.Get(key)
	if err != nil {
		return err
	}
	rec := new(ConnectionRequestRec)
	err = json.Unmarshal(data, rec)
	if err != nil {
		return err
	}

	if rec.State >= state {
		return fmt.Errorf("error state with id:%s", id)
	}

	rec.State = state
	bts, err := json.Marshal(rec)
	if err != nil {
		return err
	}
	return s.store.Put(key, bts)
}
