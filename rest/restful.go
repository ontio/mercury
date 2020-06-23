package rest

import (
	"net/http"

	"git.ont.io/ontid/otf/message"
	"git.ont.io/ontid/otf/middleware"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
)

func Invite(c *gin.Context) {
	resp := Gin{C: c}
	invite := &Invitation{}
	err := c.Bind(invite)
	if err != nil {
		middleware.Log.Errorf("Invite err:%s", err)
	}
	data, err := SendMsg(message.Invitation, structs.Map(invite))
	if err != nil {
		middleware.Log.Errorf("Invite err:%s", err)
	}
	resp.Response(http.StatusOK, 0, data)
}

func Connect(c *gin.Context) {
	resp := Gin{C: c}
	req := &ConnectionRequest{}
	err := c.Bind(req)
	if err != nil {
		middleware.Log.Errorf("ConnectionRequest err:%s", err)
	}
	data, err := SendMsg(message.ConnectionRequest, structs.Map(req))
	if err != nil {
		middleware.Log.Errorf("connect err:%s")
	}
	resp.Response(http.StatusOK, 0, data)
}

func ConnectAck(c *gin.Context) {
	resp := Gin{C: c}
	req := &ConnectionACK{}
	err := c.Bind(req)
	if err != nil {
		middleware.Log.Errorf("ConnectAck err:%s", err)
	}
	data, err := SendMsg(message.ConnectionACK, structs.Map(req))
	if err != nil {
		middleware.Log.Errorf("ConnectAck msg type:%d,err:%s", message.ConnectionACK, err)
	}
	resp.Response(http.StatusOK, 0, data)
}

func ProposalCredentialReq(c *gin.Context) {
	resp := Gin{C: c}
	req := &ProposalCredential{}
	err := c.Bind(req)
	if err != nil {
		middleware.Log.Errorf("ProposalCredential err:%s", err)
	}
	data, err := SendMsg(message.ConnectionACK, structs.Map(req))
	if err != nil {
		middleware.Log.Errorf("ProposalCredentialReq msg type:%d,err:%s", message.ConnectionACK, err)
	}
	resp.Response(http.StatusOK, 0, data)
}

func SendCredential(c *gin.Context) {
	resp := Gin{C: c}
	req := &RequestCredential{}
	err := c.Bind(req)
	if err != nil {
		middleware.Log.Errorf("ProposalCredential err:%s", err)
	}
	data, err := SendMsg(message.RequestCredential, structs.Map(req))
	if err != nil {
		middleware.Log.Errorf("ProposalCredential msg type:%d,err:%s", message.RequestCredential, err)
	}
	resp.Response(http.StatusOK, 0, data)
}

func IssueCredential(c *gin.Context) {
	resp := Gin{C: c}
	req := &OfferCredential{}
	err := c.Bind(req)
	if err != nil {
		middleware.Log.Errorf("OfferCredential err:%s", err)
	}
	data, err := SendMsg(message.IssueCredential, structs.Map(req))
	if err != nil {
		middleware.Log.Errorf("OfferCredential msg type:%d,err:%s", message.IssueCredential, err)
	}
	resp.Response(http.StatusOK, 0, data)
}

func CredentialAckInfo(c *gin.Context) {
	resp := Gin{C: c}
	req := &CredentialACK{}
	err := c.Bind(req)
	if err != nil {
		middleware.Log.Errorf("CredentialAck err:%s", err)
	}
	data, err := SendMsg(message.CredentialACK, structs.Map(req))
	if err != nil {
		middleware.Log.Errorf("CredentialAck msg type:%d,err:%s", message.CredentialACK, err)
	}
	resp.Response(http.StatusOK, 0, data)
}

func RequestProof(c *gin.Context) {
	resp := Gin{C: c}
	req := &RequestCredential{}
	err := c.Bind(req)
	if err != nil {
		middleware.Log.Errorf("RequestCredential err:%s", err)
	}
	data, err := SendMsg(message.RequestPresentation, structs.Map(req))
	if err != nil {
		middleware.Log.Errorf("RequestCredential msg type:%d,err:%s", message.RequestCredential, err)
	}
	resp.Response(http.StatusOK, 0, data)
}

func PresentProof(c *gin.Context) {
	resp := Gin{C: c}
	req := &Presentation{}
	err := c.Bind(req)
	if err != nil {
		middleware.Log.Errorf("Presentation err:%s", err)
	}
	data, err := SendMsg(message.Presentation, structs.Map(req))
	if err != nil {
		middleware.Log.Errorf("Presentation msg type:%d,err:%s", message.Presentation, err)
	}
	resp.Response(http.StatusOK, 0, data)
}

func PresentationACKInfo(c *gin.Context) {
	resp := Gin{C: c}
	req := &PresentationACK{}
	err := c.Bind(req)
	if err != nil {
		middleware.Log.Errorf("PresentationACK err:%s", err)
	}
	data, err := SendMsg(message.PresentationACK, structs.Map(req))
	if err != nil {
		middleware.Log.Errorf("PresentationACK msg type:%d,err:%s", message.PresentationACK, err)
	}
	resp.Response(http.StatusOK, 0, data)
}
