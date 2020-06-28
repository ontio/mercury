package rest

import (
	"git.ont.io/ontid/otf/config"
	"git.ont.io/ontid/otf/message"
	"git.ont.io/ontid/otf/middleware"
	"git.ont.io/ontid/otf/service"
	"git.ont.io/ontid/otf/store"
	"github.com/gin-gonic/gin"
	sdk "github.com/ontio/ontology-go-sdk"
)

var (
	Svr *service.Service
)

func NewService(acct *sdk.Account, cfg *config.Cfg, db store.Store) {
	Svr = service.NewService()
	Svr.RegisterController(service.NewSyscontroller(acct, cfg))
	Svr.RegisterController(service.NewCustomcontroller())
}

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.LoggerToFile())
	r.Use(gin.Recovery())
	v := r.Group("/api/v1")
	{
		v.POST("/invitation", Invite)
		v.POST("/connectionrequest", ConnectRequest)
		v.POST("/connectionresponse", ConnectResponse)
		v.POST("/connectionack", ConnectAck)
		v.POST("/proposalcredential", ProposalCredentialReq)
		v.POST("/sendcredential", SendCredential)
		v.POST("/issuecredentail", IssueCredential)
		v.POST("/credentialack", CredentialAckInfo)
		v.POST("/requestproof", RequestProof)
		v.POST("/presentproof", PresentProof)
		v.POST("/presentationack", PresentationACKInfo)
		v.POST("/sendgeneralmsg", SendGeneralMsg)
	}
	return r
}

func SendMsg(msgType message.MessageType, data map[string]interface{}) (interface{}, error) {
	msg := message.Message{MessageType: msgType, Content: data}
	resp, err := Svr.Serv(msg)
	if err != nil {
		middleware.Log.Errorf("err:%s", err)
		return nil, err
	}
	return resp, nil
}
