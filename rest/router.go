package rest

import (
	"git.ont.io/ontid/otf/common/config"
	"git.ont.io/ontid/otf/common/log"
	"git.ont.io/ontid/otf/common/message"
	"git.ont.io/ontid/otf/service"
	"git.ont.io/ontid/otf/service/controller"
	"git.ont.io/ontid/otf/store"
	"git.ont.io/ontid/otf/utils"
	"git.ont.io/ontid/otf/vdri"
	"github.com/gin-gonic/gin"
	sdk "github.com/ontio/ontology-go-sdk"
)

var (
	Svr           *service.Service
	EnablePackage bool
)

func NewService(acct *sdk.Account, cfg *config.Cfg, db store.Store, msgSvr *service.MsgService, v vdri.VDRI, ontSdk *sdk.OntologySdk) {
	Svr = service.NewService(ontSdk, acct)
	Svr.RegisterController(controller.NewSyscontroller(acct, cfg, db, msgSvr))
	Svr.RegisterController(controller.NewCredentialController(acct, cfg, db, msgSvr, v))
	Svr.RegisterController(controller.NewPresentationController(acct, cfg, db, msgSvr, v))
}

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Use(gin.Recovery())
	v := r.Group(utils.Group_Api_V1)
	{
		v.POST(utils.Invite_Api, Invite)
		v.POST(utils.ConnectRequest_Api, ConnectRequest)
		v.POST(utils.ConnectResponse_Api, ConnectResponse)
		v.POST(utils.ConnectAck_Api, ConnectAck)
		v.POST(utils.SendDisconnect_Api, SendDisconnect)
		v.POST(utils.Disconnect_Api, Disconnect)

		v.POST(utils.SendProposalCredentialReq_Api, SendProposalCredentialReq)
		v.POST(utils.OfferCredential_Api, OfferCredential)
		v.POST(utils.ProposalCredentialReq_Api, ProposalCredentialReq)
		v.POST(utils.SendRequestCredential_Api, SendRequestCredential)
		v.POST(utils.RequestCredential_Api, RequestCredential)
		v.POST(utils.IssueCredential_Api, IssueCredential)
		v.POST(utils.CredentialAckInfo_Api, CredentialAckInfo)

		v.POST(utils.SendRequestPresentation_Api, SendRequestPresentation)
		v.POST(utils.RequestPresentation_Api, RequestPresentation)
		v.POST(utils.Presentation_Api, Presentation)
		v.POST(utils.PresentationAckInfo_Api, PresentationAckInfo)

		v.POST(utils.SendGeneralMsg_Api, SendGeneralMsg)
		v.POST(utils.ReceiveGeneralMsg_Api, ReceiveGeneralMsg)
		v.POST(utils.QueryGeneralMsg_Api, QueryGeneralMsg)
		v.POST(utils.QueryCredential_Api, QueryCredential)
		v.POST(utils.QueryPresentation_Api, QueryPresentation)
	}
	return r
}

func SendMsg(msgType message.MessageType, data interface{}) (interface{}, error) {
	msg := message.Message{MessageType: msgType, Content: data}
	resp, err := Svr.Serv(msg)
	if err != nil {
		log.Errorf("err:%s", err)
		return nil, err
	}
	sendMsg, err := resp.GetMessage()
	if err != nil {
		return nil, err
	}
	return sendMsg.Content, nil
}
