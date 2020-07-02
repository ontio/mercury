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
	RequestPresentationKey = "RequestPresentation"
	PresentationKey        = "Presentation"
)

type PresentationController struct {
	account *sdk.Account
	did     vdri.Did
	cfg     *config.Cfg
	store   store.Store
	msgsvr  *service.MsgService
	vdri    vdri.VDRI
}

func NewPresentationController(acct *sdk.Account, cfg *config.Cfg, db store.Store, msgsvr *service.MsgService, did vdri.Did, v vdri.VDRI) PresentationController {
	p := PresentationController{
		account: acct,
		did:     did,
		cfg:     cfg,
		store:   db,
		msgsvr:  msgsvr,
		vdri:    v,
	}
	p.Initiate(nil)
	return p

}

func (p PresentationController) Initiate(param service.ParameterInf) error {
	middleware.Log.Infof("%s Initiate", p.Name())
	//todo add logic
	return nil
}

func (p PresentationController) Name() string {
	return "CredentialController"
}

func (s PresentationController) Shutdown() error {
	middleware.Log.Infof("%s shutdown\n", s.Name())
	return nil
}

func (p PresentationController) Process(msg message.Message) (service.ControllerResp, error) {
	middleware.Log.Infof("%s Process:%v\n", p.Name(), msg)
	//todo add logic
	switch msg.MessageType {
	case message.SendRequestPresentationType:
		middleware.Log.Infof("resolve SendPresentationType")
		req := msg.Content.(*message.RequestPresentation)

		outMsg := service.OutboundMsg{
			Msg: message.Message{
				MessageType: message.RequestPresentationType,
				Content:     req,
			},
			Conn: req.Connection,
		}
		err := p.msgsvr.HandleOutBound(outMsg)
		if err != nil {
			middleware.Log.Errorf("error on HandleOutBound :%s", err.Error())
			return nil, err
		}

	case message.RequestPresentationType:
		middleware.Log.Infof("resolve RequestPresentationType")
		req := msg.Content.(*message.RequestPresentation)

		presentation, err := p.vdri.PresentProof(req)
		if err != nil {
			middleware.Log.Errorf("errors on PresentProof :%s", err.Error())
			return nil, err
		}

		err = p.SaveRequestPresentation(req.Id, *req)
		if err != nil {
			middleware.Log.Errorf("error on SaveRequestPresentation:%s", err.Error())
			return nil, err
		}

		outMsg := service.OutboundMsg{
			Msg: message.Message{
				MessageType: message.PresentationType,
				Content:     presentation,
			},
			Conn: presentation.Connection,
		}
		err = p.msgsvr.HandleOutBound(outMsg)
		if err != nil {
			middleware.Log.Errorf("error on HandleOutBound:%s", err.Error())
			return nil, err
		}
	case message.PresentationType:
		middleware.Log.Infof("resolve RequestPresentationType")
		req := msg.Content.(*message.Presentation)

		err := p.SavePresentation(req.Thread.ID, *req)
		if err != nil {
			return nil, err
		}
		ack := new(message.PresentationACK)
		ack.Id = utils.GenUUID()
		ack.Thread = req.Thread
		ack.Connection = service.ReverseConnection(req.Connection)
		ack.Type = vdri.PresentationACKSpec
		ack.Status = ACK_SUCCEED

		outMsg := service.OutboundMsg{
			Msg: message.Message{
				MessageType: message.PresentationACKType,
				Content:     ack,
			},
			Conn: ack.Connection,
		}
		err = p.msgsvr.HandleOutBound(outMsg)
		if err != nil {
			middleware.Log.Errorf("error on HandleOutBound:%s", err.Error())
			return nil, err
		}

	case message.PresentationACKType:
		middleware.Log.Infof("resolve PresentationACKType")
		req := msg.Content.(*message.PresentationACK)

		err := p.UpdateRequestPresentaion(req.Thread.ID, service.RequestPresentationReceived)
		if err != nil {
			return nil, err
		}
		middleware.Log.Infof("ack received")

	default:
		return service.Skipmessage(msg)
	}

	return nil, nil

}

func (p PresentationController) SaveRequestPresentation(id string, rr message.RequestPresentation) error {
	key := []byte(fmt.Sprintf("%s_%s", RequestPresentationKey, id))
	b, err := p.store.Has(key)
	if err != nil {
		return err
	}
	if b {
		return fmt.Errorf("ReqeustPresentation id:%s,all ready exist", id)
	}

	rec := new(service.RequestPresentationRec)
	rec.RerquestPrentation = rr
	rec.RequesterDID = rr.Connection.MyDid
	rec.State = service.RequestPresentationReceived

	data, err := json.Marshal(rec)
	if err != nil {
		return err
	}

	return p.store.Put(key, data)
}

func (p PresentationController) UpdateRequestPresentaion(id string, state service.RequestPresentationState) error {
	key := []byte(fmt.Sprintf("%s_%s", RequestPresentationKey, id))
	data, err := p.store.Get(key)
	if err != nil {
		return err
	}
	rec := new(service.RequestPresentationRec)
	err = json.Unmarshal(data, rec)
	if err != nil {
		return err
	}
	if rec.State <= state {
		return fmt.Errorf("request presentation id:%s state invalid", id)
	}

	rec.State = state
	data, err = json.Marshal(rec)
	if err != nil {
		return err
	}
	return p.store.Put(key, data)
}

func (p PresentationController) SavePresentation(id string, pr message.Presentation) error {
	key := []byte(fmt.Sprintf("%s_%s", PresentationKey, id))
	b, err := p.store.Has(key)
	if err != nil {
		return err
	}
	if b {
		return fmt.Errorf("ReqeustPresentation id:%s,all ready exist", id)
	}

	rec := new(service.PresentationRec)
	rec.Presentation = pr
	rec.OwnerDID = pr.Connection.TheirDid
	rec.Timestamp = time.Now()

	data, err := json.Marshal(rec)
	if err != nil {
		return err
	}

	return p.store.Put(key, data)
}
