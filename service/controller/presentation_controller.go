package controller

import (
	"encoding/json"
	"fmt"
	"git.ont.io/ontid/otf/utils"
	"time"

	"git.ont.io/ontid/otf/common/config"
	"git.ont.io/ontid/otf/common/log"
	"git.ont.io/ontid/otf/common/message"
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
	cfg     *config.Cfg
	store   store.Store
	msgsvr  *service.MsgService
	vdri    vdri.VDRI
}

func NewPresentationController(acct *sdk.Account, cfg *config.Cfg, db store.Store, msgsvr *service.MsgService, v vdri.VDRI) PresentationController {
	p := PresentationController{
		account: acct,
		cfg:     cfg,
		store:   db,
		msgsvr:  msgsvr,
		vdri:    v,
	}
	err := p.Initiate(nil)
	if err != nil {
		panic(err)
	}
	return p

}

func (p PresentationController) Initiate(param service.ParameterInf) error {
	log.Infof("%s Initiate", p.Name())
	//todo add logic
	return nil
}

func (p PresentationController) Name() string {
	return "CredentialController"
}

func (p PresentationController) Shutdown() error {
	log.Infof("%s shutdown\n", p.Name())
	return nil
}

func (p PresentationController) Process(msg message.Message) (service.ControllerResp, error) {
	log.Infof("%s Process:%v\n", p.Name(), msg)
	switch msg.MessageType {
	case message.SendRequestPresentationType:
		log.Infof("resolve SendPresentationType")
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
			log.Errorf("error on HandleOutBound :%s", err.Error())
			return nil, err
		}

	case message.RequestPresentationType:
		log.Infof("resolve RequestPresentationType")
		req := msg.Content.(*message.RequestPresentation)
		err := utils.CheckConnection(req.Connection.TheirDid, req.Connection.MyDid, p.store)
		if err != nil {
			log.Infof("no connect found with did:%s", req.Connection.MyDid)
			return nil, err
		}
		presentation, err := p.vdri.PresentProof(req, p.store)
		if err != nil {
			log.Errorf("errors on PresentProof :%s", err.Error())
			return nil, err
		}

		err = p.SaveRequestPresentation(req.Connection.MyDid, req.Id, *req)
		if err != nil {
			log.Errorf("error on SaveRequestPresentation:%s", err.Error())
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
			log.Errorf("error on HandleOutBound:%s", err.Error())
			return nil, err
		}
	case message.PresentationType:
		log.Infof("resolve RequestPresentationType")
		req := msg.Content.(*message.Presentation)
		err := utils.CheckConnection(req.Connection.TheirDid, req.Connection.MyDid, p.store)
		if err != nil {
			log.Infof("no connect found with did:%s", req.Connection.MyDid)
			return nil, err
		}
		err = p.SavePresentation(req.Connection.TheirDid, req.Thread.ID, *req)
		if err != nil {
			return nil, err
		}
		ack := new(message.PresentationACK)
		ack.Id = utils.GenUUID()
		ack.Thread = req.Thread
		ack.Connection = service.ReverseConnection(req.Connection)
		ack.Type = vdri.PresentationACKSpec
		ack.Status = utils.ACK_SUCCEED

		outMsg := service.OutboundMsg{
			Msg: message.Message{
				MessageType: message.PresentationACKType,
				Content:     ack,
			},
			Conn: ack.Connection,
		}
		err = p.msgsvr.HandleOutBound(outMsg)
		if err != nil {
			log.Errorf("error on HandleOutBound:%s", err.Error())
			return nil, err
		}

	case message.PresentationACKType:
		log.Infof("resolve PresentationACKType")
		req := msg.Content.(*message.PresentationACK)
		err := utils.CheckConnection(req.Connection.TheirDid, req.Connection.MyDid, p.store)
		if err != nil {
			log.Infof("no connect found with did:%s", req.Connection.MyDid)
			return nil, err
		}
		err = p.UpdateRequestPresentaion(req.Connection.MyDid, req.Thread.ID, message.RequestPresentationReceived)
		if err != nil {
			return nil, err
		}
		log.Infof("ack received")

	case message.QueryPresentationType:
		log.Infof("resolve QueryPresentationType")
		req := msg.Content.(*message.QueryPresentationRequest)

		rec, err := p.QueryPresentation(req.DId, req.Id)
		if err != nil {
			log.Errorf("error on QueryPresentationType:%s", err.Error())
			return nil, err
		}
		queryRespons := new(message.QueryPresentationResponse)
		queryRespons.Formats = rec.Formats
		queryRespons.PresentationAttach = rec.PresentationAttach

		return service.ServiceResponse{
			Message: queryRespons,
		}, nil

	default:
		return service.Skipmessage(msg)
	}

	return nil, nil

}

func (p PresentationController) SaveRequestPresentation(did, id string, rr message.RequestPresentation) error {
	key := []byte(fmt.Sprintf("%s_%s_%s", RequestPresentationKey, did, id))
	b, err := p.store.Has(key)
	if err != nil {
		return err
	}
	if b {
		return fmt.Errorf("ReqeustPresentation id:%s,all ready exist", id)
	}

	rec := new(message.RequestPresentationRec)
	rec.RerquestPrentation = rr
	rec.RequesterDID = rr.Connection.MyDid
	rec.State = message.RequestPresentationReceived

	data, err := json.Marshal(rec)
	if err != nil {
		return err
	}

	return p.store.Put(key, data)
}

func (p PresentationController) UpdateRequestPresentaion(did, id string, state message.RequestPresentationState) error {
	key := []byte(fmt.Sprintf("%s_%s_%s", RequestPresentationKey, did, id))
	data, err := p.store.Get(key)
	if err != nil {
		return err
	}
	rec := new(message.RequestPresentationRec)
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

func (p PresentationController) SavePresentation(did, id string, pr message.Presentation) error {
	key := []byte(fmt.Sprintf("%s_%s_%s", PresentationKey, did, id))
	b, err := p.store.Has(key)
	if err != nil {
		return err
	}
	if b {
		return fmt.Errorf("ReqeustPresentation id:%s,all ready exist", id)
	}

	rec := new(message.PresentationRec)
	rec.Presentation = pr
	rec.OwnerDID = pr.Connection.TheirDid
	rec.Timestamp = time.Now()

	data, err := json.Marshal(rec)
	if err != nil {
		return err
	}

	return p.store.Put(key, data)
}

func (p PresentationController) QueryPresentation(did, id string) (message.Presentation, error) {
	key := []byte(fmt.Sprintf("%s_%s_%s", PresentationKey, did, id))
	data, err := p.store.Get(key)
	if err != nil {
		return message.Presentation{}, err
	}
	rec := new(message.PresentationRec)
	err = json.Unmarshal(data, rec)
	if err != nil {
		return message.Presentation{}, err
	}
	return rec.Presentation, nil
}
