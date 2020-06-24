package rest

import (
	"git.ont.io/ontid/otf/config"
	"git.ont.io/ontid/otf/message"
	"git.ont.io/ontid/otf/middleware"
	"git.ont.io/ontid/otf/service"
	"github.com/gin-gonic/gin"
	sdk "github.com/ontio/ontology-go-sdk"

)


var (
	Svr *service.Service
)

func NewService(acct *sdk.Account,cfg *config.Cfg ) {
	Svr := service.NewService()
	Svr.RegisterController(service.NewSyscontroller(acct,cfg))
	Svr.RegisterController(service.NewCustomcontroller())
}

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.LoggerToFile())
	r.Use(gin.Recovery())
	v1 := r.Group("/api/v1")
	{
		v1.POST("/invitation", Invite)
		v1.POST("/connection", Connect)
		v1.POST("/connectionack", ConnectAck)
		v1.POST("/proposalcredential", ProposalCredentialReq)
		v1.POST("/sendcredential", SendCredential)
		v1.POST("/issuecredentail", IssueCredential)
		v1.POST("/credentialack", CredentialAckInfo)
		v1.POST("/requestproof", RequestProof)
		v1.POST("/presentproof", PresentProof)
		v1.POST("/presentationack", PresentationACKInfo)
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
