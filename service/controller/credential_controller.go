package controller

import (
	"encoding/json"
	"fmt"
	"time"

	"git.ont.io/ontid/otf/config"
	"git.ont.io/ontid/otf/message"
	"git.ont.io/ontid/otf/service"
	"git.ont.io/ontid/otf/store"
	"git.ont.io/ontid/otf/vdri"
	"github.com/google/uuid"
	sdk "github.com/ontio/ontology-go-sdk"
)

const (
	ProposalCredentialSpec = "spec/issue-credential/" + Version + "/propose-credential"
	OfferCredentialSpec    = "spec/issue-credential/" + Version + "/offer-credential"
	RequestCredentialSpec  = "spec/issue-credential/" + Version + "/request-credential"
	IssueCredentialSpec    = "spec/issue-credential/" + Version + "/issue-credential"
	CredentialACKSpec      = "spec/issue-credential/" + Version + "/ack"

	CredentialKey        = "Credential"
	RequestCredentialKey = "RequestCredential"
	OfferCredentialKey   = "OfferCredential"
)

type CredentialController struct {
	account *sdk.Account
	did     vdri.Did
	cfg     *config.Cfg
	store   store.Store
	msgsvr  *service.MsgService
}

func NewCredentialController(acct *sdk.Account, cfg *config.Cfg, db store.Store, msgsvr *service.MsgService, did vdri.Did) CredentialController {
	s := CredentialController{
		account: acct,
		did:     did,
		cfg:     cfg,
		store:   db,
		msgsvr:  msgsvr,
	}
	s.Initiate(nil)
	return s

}

func (s CredentialController) Name() string {
	return "CredentialController"
}

func (s CredentialController) Initiate(param service.ParameterInf) error {
	fmt.Printf("%s Initiate\n", s.Name())
	//todo add logic
	return nil
}

func (s CredentialController) Process(msg message.Message) (service.ControllerResp, error) {
	fmt.Printf("%s Process:%v\n", s.Name(), msg)
	//todo add logic
	switch msg.MessageType {
	case message.SendProposalCredentialType:
		fmt.Printf("resolve SendProposalCredentialType")
		req := msg.Content.(*message.ProposalCredential)

		outMsg := service.OutboundMsg{
			Msg: message.Message{
				MessageType: message.ProposalCredentialType,
				Content:     req,
			},
			Conn: req.Connection,
		}
		err := s.msgsvr.HandleOutBound(outMsg)
		if err != nil {
			return nil, err
		}

	case message.ProposalCredentialType:
		fmt.Printf("resolve ProposalCredentialType")
		req := msg.Content.(*message.ProposalCredential)
		//todo deal with the proposal, do we need store the proposal???
		fmt.Printf("proposal is %v\n", req)

		//for sample only
		offer := new(message.OfferCredential)
		offer.Type = OfferCredentialSpec
		offer.Id = uuid.New().String()
		offer.Connection = service.ReverseConnection(req.Connection)
		offer.CredentialPreview = message.CredentialPreview{Type: "sample", Attributre: []message.Attributre{message.Attributre{
			Name:     "name1",
			MimeType: "json",
			Value:    "{abc}",
		}}}
		offer.Thread = message.Thread{
			ID: req.Id,
		}

		outerMsg := service.OutboundMsg{
			Msg: message.Message{
				MessageType: message.OfferCredentialType,
				Content:     offer,
			},
			Conn: offer.Connection,
		}

		err := s.msgsvr.HandleOutBound(outerMsg)
		if err != nil {
			fmt.Printf("error on HandleOutBound :%s\n", err.Error())
			return nil, err
		}

	case message.OfferCredentialType:
		fmt.Printf("resolve ProposalCredentialType")
		req := msg.Content.(*message.OfferCredential)
		fmt.Printf("req:%v\n", req)
		//todo save the offer in store
		err := s.SaveOfferCredential(req.Thread.ID, req)
		if err != nil {
			return nil, err
		}
		//

	case message.SendRequestCredentialType:
		fmt.Printf("resolve SendRequestCredentialType")
		req := msg.Content.(*message.RequestCredential)
		outMsg := service.OutboundMsg{
			Msg: message.Message{
				MessageType: message.RequestCredentialType,
				Content:     req,
			},
			Conn: req.Connection,
		}
		err := s.msgsvr.HandleOutBound(outMsg)
		if err != nil {
			return nil, err
		}

	case message.RequestCredentialType:
		fmt.Printf("resolve RequestCredentialType")
		req := msg.Content.(*message.RequestCredential)

		//todo deal with the request
		fmt.Printf("request is %v\n", req)

		//todo store the request
		err := s.SaveRequestCredential(req.Id, *req)
		if err != nil {
			fmt.Printf("error on SaveRequestCredential:%s\n", err.Error())
			return nil, err
		}
		credential := message.IssueCredential{
			Type:              IssueCredentialSpec,
			Id:                uuid.New().String(),
			Comment:           "",
			Formats:           nil,
			CredentialsAttach: nil,
			Connection:        service.ReverseConnection(req.Connection),
			Thread: message.Thread{
				ID: req.Id,
			},
		}

		outMsg := service.OutboundMsg{
			Msg: message.Message{
				MessageType: message.IssueCredentialType,
				Content:     credential,
			},
			Conn: credential.Connection,
		}

		err = s.msgsvr.HandleOutBound(outMsg)
		if err != nil {
			fmt.Printf("error on HandleOutBound:%s\n", err.Error())
			return nil, err
		}

	case message.IssueCredentialType:
		fmt.Printf("resolve IssueCredentialType")
		req := msg.Content.(*message.IssueCredential)

		//store the credential
		err := s.SaveCredential(req.Thread.ID, *req)
		if err != nil {
			fmt.Printf("error on SaveCredential:%s\n", err.Error())
			return nil, err
		}

		ack := message.CredentialACK{
			Type: CredentialACKSpec,
			Id:   uuid.New().String(),
			Thread: message.Thread{
				ID: req.Thread.ID,
			},
			Status:     ACK_SUCCEED,
			Connection: service.ReverseConnection(req.Connection),
		}

		outmsg := service.OutboundMsg{
			Msg: message.Message{
				MessageType: message.CredentialACKType,
				Content:     ack,
			},
			Conn: ack.Connection,
		}
		err = s.msgsvr.HandleOutBound(outmsg)
		if err != nil {
			fmt.Printf("error on SaveCredential:%s\n", err.Error())
			return nil, err
		}

	case message.CredentialACKType:
		fmt.Printf("resolve IssueCredentialType")
		req := msg.Content.(*message.CredentialACK)
		reqid := req.Thread.ID

		err := s.UpdateRequestCredential(reqid, service.RequestCredentialResolved)
		if err != nil {
			fmt.Printf("error on UpdateRequestCredential:%s\n", err.Error())
			return nil, err
		}

	default:
		return service.Skipmessage(msg)
	}

	return nil, nil
}
func (s CredentialController) Shutdown() error {
	fmt.Printf("%s shutdown\n", s.Name())
	return nil
}

func (s CredentialController) SaveOfferCredential(id string, propsal *message.OfferCredential) error {
	key := []byte(fmt.Sprintf("%s_%s", OfferCredentialKey, id))
	b, err := s.store.Has(key)
	if err != nil {
		return err
	}
	if b {
		return fmt.Errorf("id:%s already exist\n", id)
	}

	data, err := json.Marshal(propsal)
	if err != nil {
		return err
	}
	return s.store.Put(key, data)
	return nil
}

func (s CredentialController) SaveCredential(id string, credential message.IssueCredential) error {
	key := []byte(fmt.Sprintf("%s_%s", CredentialKey, id))
	b, err := s.store.Has(key)
	if err != nil {
		return err
	}
	if b {
		return fmt.Errorf("id:%s already exist\n", id)
	}

	rec := service.CredentialRec{
		OwnerDID:   credential.Connection.TheirServiceId,
		Credential: credential,
		Timestamp:  time.Now(),
	}
	data, err := json.Marshal(rec)
	if err != nil {
		return err
	}
	return s.store.Put(key, data)
}

func (s CredentialController) SaveRequestCredential(id string, requestCredential message.RequestCredential) error {
	key := []byte(fmt.Sprintf("%s_%s", RequestCredentialKey, id))
	b, err := s.store.Has(key)
	if err != nil {
		return err
	}
	if b {
		return fmt.Errorf("id:%s already exist\n", id)
	}

	rec := service.RequestCredentialRec{
		RequesterDID:      requestCredential.Connection.MyDid,
		RequestCredential: requestCredential,
		State:             service.RequestCredentialReceived,
	}
	data, err := json.Marshal(rec)
	if err != nil {
		return err
	}
	return s.store.Put(key, data)
}

func (s CredentialController) UpdateRequestCredential(id string, state service.RequestCredentialState) error {
	key := []byte(fmt.Sprintf("%s_%s", RequestCredentialKey, id))
	data, err := s.store.Get(key)
	if err != nil {
		return err
	}

	rec := new(service.RequestCredentialRec)
	err = json.Unmarshal(data, rec)
	if err != nil {
		return err
	}
	if rec.State >= state {
		return fmt.Errorf("UpdateRequestCredential id :%s state invalid\n")
	}
	rec.State = state
	data, err = json.Marshal(rec)
	if err != nil {
		return err
	}
	return s.store.Put(key, data)

}
