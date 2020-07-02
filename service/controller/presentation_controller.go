package controller

import (
	"fmt"
	"git.ont.io/ontid/otf/config"
	"git.ont.io/ontid/otf/message"
	"git.ont.io/ontid/otf/service"
	"git.ont.io/ontid/otf/store"
	"git.ont.io/ontid/otf/vdri"
	ontdid "git.ont.io/ontid/otf/vdri/ontdid"
	sdk "github.com/ontio/ontology-go-sdk"
)

type PresentationController struct {
	account *sdk.Account
	did     vdri.Did
	cfg     *config.Cfg
	store   store.Store
	msgsvr  *service.MsgService
}

func NewPresentationController(acct *sdk.Account, cfg *config.Cfg, db store.Store, msgsvr *service.MsgService) PresentationController {
	did := ontdid.NewOntDID(cfg, acct)
	p := PresentationController{
		account: acct,
		did:     did,
		cfg:     cfg,
		store:   db,
		msgsvr:  msgsvr,
	}
	p.Initiate(nil)
	return p

}

func (p PresentationController) Initiate(param service.ParameterInf) error {
	fmt.Printf("%s Initiate\n", p.Name())
	//todo add logic
	return nil
}

func (p PresentationController) Name() string {
	return "CredentialController"
}

func (p PresentationController) Process(msg message.Message) (service.ControllerResp, error) {
	fmt.Printf("%s Process:%v\n", p.Name(), msg)
	//todo add logic
	switch msg.MessageType {
	case message.SendPresentationType:
		fmt.Printf("resolve SendPresentationType")
		req := msg.Content.(*message.RequestPresentation)

		outMsg := service.OutboundMsg{
			Msg: message.Message{
				MessageType: message.RequestPresentationType,
				Content:     req,
			},
			Conn: req.Connection,
		}
		err := p.msgsvr.HandleOutBound(outMsg)
		if err != nil{
			return nil,err
		}

	case message.RequestPresentationType:

	case message.PresentationType:
	case message.PresentationACKType:

	default:
		return service.Skipmessage(msg)

	}

	return nil, nil

}
