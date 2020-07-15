package controller

import (
	"encoding/json"
	"fmt"
	"git.ont.io/ontid/otf/middleware"
	"git.ont.io/ontid/otf/utils"
	"time"

	"git.ont.io/ontid/otf/config"
	"git.ont.io/ontid/otf/message"
	"git.ont.io/ontid/otf/service"
	"git.ont.io/ontid/otf/store"
	"git.ont.io/ontid/otf/vdri"
	sdk "github.com/ontio/ontology-go-sdk"
)

const (
	CredentialKey        = "Credential"
	RequestCredentialKey = "RequestCredential"
	OfferCredentialKey   = "OfferCredential"
)

type CredentialController struct {
	account *sdk.Account
	cfg     *config.Cfg
	store   store.Store
	msgsvr  *service.MsgService
	vdri    vdri.VDRI
}

func NewCredentialController(acct *sdk.Account, cfg *config.Cfg, db store.Store, msgsvr *service.MsgService, v vdri.VDRI) CredentialController {
	s := CredentialController{
		account: acct,
		cfg:     cfg,
		store:   db,
		msgsvr:  msgsvr,
		vdri:    v,
	}
	err := s.Initiate(nil)
	if err != nil {
		panic(err)
	}
	return s

}

func (s CredentialController) Name() string {
	return "CredentialController"
}

func (s CredentialController) Initiate(param service.ParameterInf) error {
	middleware.Log.Infof("%s Initiate", s.Name())
	//todo add logic
	return nil
}

func (s CredentialController) Process(msg message.Message) (service.ControllerResp, error) {
	middleware.Log.Infof("%s Process:%v", s.Name(), msg)
	//todo add logic
	switch msg.MessageType {
	case message.SendProposalCredentialType:
		middleware.Log.Infof("resolve SendProposalCredentialType")
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
			middleware.Log.Errorf("error on HandleOutBound:%s", err.Error())
			return nil, err
		}

	case message.ProposalCredentialType:
		middleware.Log.Infof("resolve ProposalCredentialType")
		req := msg.Content.(*message.ProposalCredential)

		err := utils.CheckConnection(req.Connection.TheirDid, req.Connection.MyDid, s.store)
		if err != nil {
			middleware.Log.Infof("no connect found with did:%s", req.Connection.MyDid)
			return nil, err
		}

		//todo deal with the proposal, do we need store the proposal???
		middleware.Log.Infof("proposal is %v", req)

		offer, err := s.vdri.OfferCredential(req)
		if err != nil {
			middleware.Log.Errorf("error on offerCredetial")
			return nil, err
		}

		outerMsg := service.OutboundMsg{
			Msg: message.Message{
				MessageType: message.OfferCredentialType,
				Content:     offer,
			},
			Conn: offer.Connection,
		}

		err = s.msgsvr.HandleOutBound(outerMsg)
		if err != nil {
			middleware.Log.Errorf("error on HandleOutBound :%s", err.Error())
			return nil, err
		}

	case message.OfferCredentialType:
		middleware.Log.Infof("resolve ProposalCredentialType")
		req := msg.Content.(*message.OfferCredential)

		err := utils.CheckConnection(req.Connection.TheirDid, req.Connection.MyDid, s.store)
		if err != nil {
			middleware.Log.Infof("no connect found with did:%s", req.Connection.MyDid)
			return nil, err
		}
		//todo save the offer in store
		err = s.SaveOfferCredential(req.Connection.TheirDid, req.Thread.ID, req)
		if err != nil {
			middleware.Log.Errorf("error on SaveOfferCredential:%s", err.Error())
			return nil, err
		}
		//

	case message.SendRequestCredentialType:
		middleware.Log.Infof("resolve SendRequestCredentialType")
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
			middleware.Log.Errorf("error on HandleOutBound:%s", err.Error())
			return nil, err
		}

	case message.RequestCredentialType:
		middleware.Log.Infof("resolve RequestCredentialType")
		req := msg.Content.(*message.RequestCredential)
		err := utils.CheckConnection(req.Connection.TheirDid, req.Connection.MyDid, s.store)
		if err != nil {
			middleware.Log.Infof("no connect found with did:%s", req.Connection.MyDid)
			return nil, err
		}
		err = s.SaveRequestCredential(req.Connection.MyDid, req.Id, *req)
		if err != nil {
			middleware.Log.Errorf("error on SaveRequestCredential:%s\n", err.Error())
			return nil, err
		}

		credential, err := s.vdri.IssueCredential(req)
		if err != nil {
			middleware.Log.Errorf("error on IssueCredential:%s\n", err.Error())
			return nil, err
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
			middleware.Log.Errorf("error on HandleOutBound:%s\n", err.Error())
			return nil, err
		}

	case message.IssueCredentialType:
		middleware.Log.Infof("resolve IssueCredentialType")
		req := msg.Content.(*message.IssueCredential)
		err := utils.CheckConnection(req.Connection.TheirDid, req.Connection.MyDid, s.store)
		if err != nil {
			middleware.Log.Infof("no connect found with did:%s", req.Connection.MyDid)
			return nil, err
		}
		//store the credential
		err = s.SaveCredential(req.Connection.TheirDid, req.Thread.ID, *req)
		if err != nil {
			middleware.Log.Errorf("error on SaveCredential:%s\n", err.Error())
			return nil, err
		}

		ack := message.CredentialACK{
			Type: vdri.CredentialACKSpec,
			Id:   utils.GenUUID(),
			Thread: message.Thread{
				ID: req.Thread.ID,
			},
			Status:     utils.ACK_SUCCEED,
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
			middleware.Log.Errorf("error on SaveCredential:%s\n", err.Error())
			return nil, err
		}

	case message.CredentialACKType:
		middleware.Log.Infof("resolve IssueCredentialType")
		req := msg.Content.(*message.CredentialACK)
		reqid := req.Thread.ID

		err := s.UpdateRequestCredential(req.Connection.MyDid, reqid, message.RequestCredentialResolved)
		if err != nil {
			middleware.Log.Errorf("error on UpdateRequestCredential:%s\n", err.Error())
			return nil, err
		}

	case message.QueryCredentialType:
		middleware.Log.Infof("resolve QueryCredentialType")
		req := msg.Content.(*message.QueryCredentialRequest)
		fmt.Printf("did:%s,id:%s\n", req.DId, req.Id)

		rec, err := s.QueryCredential(req.DId, req.Id)
		if err != nil {
			middleware.Log.Errorf("error on QueryCredentialType:%s\n", err.Error())
			return nil, err
		}
		resp := new(service.ServiceResponse)

		queryResult := new(message.QueryCredentialResponse)
		queryResult.CredentialsAttach = rec.CredentialsAttach
		queryResult.Formats = rec.Formats

		resp.Message = queryResult
		return resp, nil

	default:
		return service.Skipmessage(msg)
	}

	return nil, nil
}
func (s CredentialController) Shutdown() error {
	middleware.Log.Infof("%s shutdown\n", s.Name())
	return nil
}

func (s CredentialController) SaveOfferCredential(did, id string, propsal *message.OfferCredential) error {
	key := []byte(fmt.Sprintf("%s_%s_%s", OfferCredentialKey, did, id))
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
}

func (s CredentialController) SaveCredential(did, id string, credential message.IssueCredential) error {
	key := []byte(fmt.Sprintf("%s_%s_%s", CredentialKey, did, id))
	b, err := s.store.Has(key)
	if err != nil {
		return err
	}
	if b {
		return fmt.Errorf("id:%s already exist\n", id)
	}

	rec := message.CredentialRec{
		OwnerDID:   credential.Connection.TheirDid,
		Credential: credential,
		Timestamp:  time.Now(),
	}
	data, err := json.Marshal(rec)
	if err != nil {
		return err
	}
	return s.store.Put(key, data)
}

func (s CredentialController) SaveRequestCredential(did, id string, requestCredential message.RequestCredential) error {
	key := []byte(fmt.Sprintf("%s_%s_%s", RequestCredentialKey, did, id))
	b, err := s.store.Has(key)
	if err != nil {
		return err
	}
	if b {
		return fmt.Errorf("id:%s already exist\n", id)
	}

	rec := message.RequestCredentialRec{
		RequesterDID:      requestCredential.Connection.MyDid,
		RequestCredential: requestCredential,
		State:             message.RequestCredentialReceived,
	}
	data, err := json.Marshal(rec)
	if err != nil {
		return err
	}
	return s.store.Put(key, data)
}

func (s CredentialController) QueryCredential(did, id string) (message.IssueCredential, error) {
	key := []byte(fmt.Sprintf("%s_%s_%s", CredentialKey, did, id))

	data, err := s.store.Get(key)
	if err != nil {
		return message.IssueCredential{}, err
	}

	rec := new(message.CredentialRec)
	err = json.Unmarshal(data, rec)
	if err != nil {
		return message.IssueCredential{}, err
	}
	return rec.Credential, nil
}

func (s CredentialController) UpdateRequestCredential(did, id string, state message.RequestCredentialState) error {
	key := []byte(fmt.Sprintf("%s_%s_%s", RequestCredentialKey, did, id))
	data, err := s.store.Get(key)
	if err != nil {
		return err
	}

	rec := new(message.RequestCredentialRec)
	err = json.Unmarshal(data, rec)
	if err != nil {
		return err
	}
	if rec.State >= state {
		return fmt.Errorf("UpdateRequestCredential id :%s state invalid\n", id)
	}
	rec.State = state
	data, err = json.Marshal(rec)
	if err != nil {
		return err
	}
	return s.store.Put(key, data)
}
