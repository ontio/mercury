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
	Version            = "1.0"
	InvitationSpec     = "spec/connections/" + Version + "/invitation"
	ConnectionRequest  = "spec/connections/" + Version + "/request"
	ConnectionResponse = "spec/connections/" + Version + "/response"
	ACK                = "spec/didcomm/" + Version + "/ack"

	InvitationKey    = "Invitation"
	ConnectionReqKey = "ConnectionReq"
	ConnectionKey    = "Connection"

	ACK_SUCCEED = "succeed"
	ACK_FAILED  = "failed"
)

type Syscontroller struct {
	account *sdk.Account
	did     did.Did
	cfg     *config.Cfg
	store   store.Store
	msgsvr  *MsgService
}

func NewSyscontroller(acct *sdk.Account, cfg *config.Cfg, db store.Store, msgsvr *MsgService) Syscontroller {
	did := did.NewOntDID(cfg, acct)
	s := Syscontroller{
		account: acct,
		did:     did,
		cfg:     cfg,
		store:   db,
		msgsvr:  msgsvr,
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
		fmt.Println("resolve invitation")
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

		return ServiceResp{
			Message: invitation,
		}, nil

	//not in use anymore
	//case message.SendConnectionRequestType:
	//	//send connection req for agent
	//	fmt.Println("resolve send connection request")
	//	//todo verify request
	//	if msg.Content == nil {
	//		return nil, fmt.Errorf("message content is nil")
	//	}
	//	cr := msg.Content.(*message.ConnectionRequest)
	//	cr.Id = uuid.New().String()
	//	err := s.SaveConnectionRequest(*cr, ConnectionRequestSent)
	//	if err != nil {
	//		return nil, err
	//	}
	//
	//	//send the connection req to target service endpoint
	//	msg.Content = cr
	//	msg.MessageType = message.ConnectionRequestType
	//	jsonbytes,err := json.Marshal(msg)
	//	if err != nil{
	//		return nil,err
	//	}
	//	msg.JsonBytes = jsonbytes
	//
	//	err = s.msgsvr.HandleOutBound(msg)
	//	if err != nil {
	//		return nil, err
	//	}
	//	middleware.Log.Infof("SendConnectionReq:%v", cr)
	//	//no need to pass incoming param
	//	return nil, nil

	case message.ConnectionRequestType:
		fmt.Println("resolve connection request")
		if msg.Content == nil {
			return nil, fmt.Errorf("message content is nil")
		}
		req := msg.Content.(*message.ConnectionRequest)
		//ivid := req.Thread.ID
		ivrc, err := s.GetInvitation(req.InvitationId)
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
		res := new(message.ConnectionResponse)
		res.Id = uuid.New().String()
		res.Thread = message.Thread{
			ID: req.Id,
		}
		//todo define the response type
		res.Type = ConnectionResponse
		//self conn
		res.Connection = message.Connection{
			Did:       ivrc.Invitation.Did,
			ServiceId: ivrc.Invitation.ServiceId,
		}

		outmsg := message.Message{
			MessageType: message.ConnectionResponseType,
			Content:     res,
		}
		err = s.msgsvr.HandleOutBound(OutboundMsg{
			Msg:  outmsg,
			Conn: req.Connection,
		})
		if err != nil {
			return nil, err
		}
		middleware.Log.Infof("ConnectionReq:%v", outmsg)
		return nil, nil

	case message.ConnectionResponseType:
		fmt.Println("resolve connection response")
		if msg.Content == nil {
			return nil, fmt.Errorf("message content is nil")
		}
		req := msg.Content.(*message.ConnectionResponse)
		connid := req.Thread.ID
		//1. update connection request to receive response state
		//err := s.UpdateConnectionRequest(connid, ConnectionResponseReceived)
		//if err != nil {
		//	return nil, err
		//}

		//2. create and save a connection object
		err := s.SaveConnection(req.Connection.Did, req.Connection.ServiceId)
		if err != nil {
			return nil, err
		}

		//3. send ACK back
		ack := message.GeneralACK{
			Type:   ACK,
			Id:     uuid.New().String(),
			Thread: message.Thread{ID: connid},
			Status: ACK_SUCCEED,
		}

		jsonbytes, err := json.Marshal(ack)
		if err != nil {
			return nil, err
		}

		outmsg := message.Message{
			MessageType: message.ConnectionACKType,
			Content:     ack,
			JsonBytes:   jsonbytes,
		}
		err = s.msgsvr.HandleOutBound(OutboundMsg{
			Msg:  outmsg,
			Conn: req.Connection,
		})
		if err != nil {
			return nil, err
		}
		return nil, nil
	case message.ConnectionACKType:
		fmt.Println("resolve ConnectionACK")
		req := msg.Content.(*message.GeneralACK)
		//1. update connection request to receive ack state
		if req.Status != ACK_SUCCEED {
			//todo remove connectionreq when failed?
			return nil, fmt.Errorf("got failed ACK ")
		}
		connid := req.Thread.ID
		err := s.UpdateConnectionRequest(connid, ConnectionACKReceived)
		if err != nil {
			return nil, err
		}
		//2. create and save a connection object
		cr, err := s.GetConnectionRequest(connid)
		if err != nil {
			return nil, err
		}

		err = s.SaveConnection(cr.ConnReq.Connection.Did, cr.ConnReq.Connection.ServiceId)
		if err != nil {
			return nil, err
		}
		return nil, nil

	case message.GeneralMsgType:
		fmt.Println("resolve GeneralMsgType")
		req := msg.Content.(message.BasicMessage)
		data, err := json.Marshal(req)
		if err != nil {
			return nil, err
		}
		fmt.Println("we got a message: %s", data)
		return nil, nil

	//for custom
	case message.ProposalCredentialType,
		message.OfferCredentialType,
		message.RequestCredentialType,
		message.IssueCredentialType,
		message.CredentialACKType:
		return skipmessage(msg)

	case message.RequestPresentationType,
		message.PresentationType,
		message.PresentationACKType:
		return skipmessage(msg)
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
	key := []byte(fmt.Sprintf("%s_%s", ConnectionReqKey, cr.Id))
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

func (s Syscontroller) GetConnectionRequest(id string) (*ConnectionRequestRec, error) {
	key := []byte(fmt.Sprintf("%s_%s", ConnectionReqKey, id))
	data, err := s.store.Get(key)
	if err != nil {
		return nil, err
	}
	cr := new(ConnectionRequestRec)
	err = json.Unmarshal(data, cr)
	if err != nil {
		return nil, err
	}
	return cr, nil
}

func (s Syscontroller) UpdateConnectionRequest(id string, state ConnectionState) error {
	key := []byte(fmt.Sprintf("%s_%s", ConnectionReqKey, id))
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

func (s Syscontroller) SaveConnection(theirDID string, serviceID string) error {
	mydid := s.did.String()

	cr := new(ConnectionRec)

	key := []byte(fmt.Sprintf("%s_%s", ConnectionKey, mydid))
	exist, err := s.store.Has(key)
	if err != nil {
		return err
	}

	if exist {
		data, err := s.store.Get(key)
		if err != nil {
			return err
		}
		err = json.Unmarshal(data, cr)
		if err != nil {
			return err
		}
		cr.Connections[fmt.Sprintf("%s_%s", theirDID, serviceID)] = message.Connection{
			Did:       theirDID,
			ServiceId: serviceID,
		}
	} else {
		cr.OwnerDID = mydid
		m := make(map[string]message.Connection)
		m[fmt.Sprintf("%s_%s", theirDID, serviceID)] = message.Connection{
			Did:       theirDID,
			ServiceId: serviceID,
		}
		cr.Connections = m
	}
	bts, err := json.Marshal(cr)
	if err != nil {
		return err
	}
	return s.store.Put(key, bts)

}
