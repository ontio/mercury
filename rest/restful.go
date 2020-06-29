package rest

import (
	"net/http"

	"git.ont.io/ontid/otf/message"
	"git.ont.io/ontid/otf/middleware"
	"git.ont.io/ontid/otf/service"
	"git.ont.io/ontid/otf/utils"
	"github.com/gin-gonic/gin"
)

func Invite(c *gin.Context) {
	resp := Gin{C: c}
	req := &message.Invitation{}
	err := c.Bind(req)
	if err != nil {
		middleware.Log.Errorf("Invite err:%s", err)
		resp.Response(http.StatusOK, 0, err.Error(), nil)
		return
	}
	data, err := SendMsg(message.InvitationType, req)
	if err != nil {
		middleware.Log.Errorf("Invite err:%s", err)
		resp.Response(http.StatusOK, 0, err.Error(), nil)
		return
	}

	jsonbytes, err := data.(service.ControllerResp).GetJsonbytes()
	if err != nil {
		middleware.Log.Errorf("Invite err:%s", err)
		resp.Response(http.StatusOK, 0, err.Error(), nil)
		return
	}

	resp.Response(http.StatusOK, 0, "", utils.Base64Encode(jsonbytes))
}

func ConnectRequest(c *gin.Context) {
	resp := Gin{C: c}
	req := &message.ConnectionRequest{}
	err := c.Bind(req)
	if err != nil {
		middleware.Log.Errorf("ConnectionRequest err:%s", err)
		resp.Response(http.StatusOK, 0, err.Error(), nil)
		return
	}
	data, err := SendMsg(message.ConnectionRequestType, req)
	if err != nil {
		middleware.Log.Errorf("connect err:%s")
		resp.Response(http.StatusOK, 0, err.Error(), nil)
		return
	}
	resp.Response(http.StatusOK, 0, "", data)
}

func ConnectResponse(c *gin.Context) {
	resp := Gin{C: c}
	req := &message.ConnectResponse{}
	err := c.Bind(req)
	if err != nil {
		middleware.Log.Errorf("ConnectResponse err:%s", err)
		resp.Response(http.StatusOK, 0, err.Error(), nil)
		return
	}
	data, err := SendMsg(message.ConnectionResponseType, req)
	if err != nil {
		middleware.Log.Errorf("connect err:%s")
		resp.Response(http.StatusOK, 0, err.Error(), nil)
		return
	}
	resp.Response(http.StatusOK, 0, "", data)
}

func ConnectAck(c *gin.Context) {
	resp := Gin{C: c}
	req := &message.GeneralACK{}
	err := c.Bind(req)
	if err != nil {
		middleware.Log.Errorf("ConnectAck err:%s", err)
		resp.Response(http.StatusOK, 0, err.Error(), nil)
		return
	}
	data, err := SendMsg(message.ConnectionACKType, req)
	if err != nil {
		middleware.Log.Errorf("ConnectAck msg type:%d,err:%s", message.ConnectionACKType, err)
		resp.Response(http.StatusOK, 0, err.Error(), nil)
		return
	}
	resp.Response(http.StatusOK, 0, "", data)
}

func ProposalCredentialReq(c *gin.Context) {
	resp := Gin{C: c}
	req := &message.ProposalCredential{}
	err := c.Bind(req)
	if err != nil {
		middleware.Log.Errorf("ProposalCredential err:%s", err)
		resp.Response(http.StatusOK, 0, err.Error(), nil)
		return
	}
	data, err := SendMsg(message.ConnectionACKType, req)
	if err != nil {
		middleware.Log.Errorf("ProposalCredentialReq msg type:%d,err:%s", message.ConnectionACKType, err)
		resp.Response(http.StatusOK, 0, err.Error(), nil)
		return
	}
	resp.Response(http.StatusOK, 0, err.Error(), data)
}

func SendCredential(c *gin.Context) {
	resp := Gin{C: c}
	req := &message.RequestCredential{}
	err := c.Bind(req)
	if err != nil {
		middleware.Log.Errorf("ProposalCredential err:%s", err)
		resp.Response(http.StatusOK, 0, err.Error(), nil)
		return
	}
	data, err := SendMsg(message.RequestCredentialType, req)
	if err != nil {
		middleware.Log.Errorf("ProposalCredential msg type:%d,err:%s", message.RequestCredentialType, err)
		resp.Response(http.StatusOK, 0, err.Error(), nil)
		return
	}
	resp.Response(http.StatusOK, 0, "", data)
}

func IssueCredential(c *gin.Context) {
	resp := Gin{C: c}
	req := &message.OfferCredential{}
	err := c.Bind(req)
	if err != nil {
		middleware.Log.Errorf("OfferCredential err:%s", err)
		resp.Response(http.StatusOK, 0, err.Error(), nil)
		return
	}
	data, err := SendMsg(message.IssueCredentialType, req)
	if err != nil {
		middleware.Log.Errorf("OfferCredential msg type:%d,err:%s", message.IssueCredentialType, err)
		resp.Response(http.StatusOK, 0, err.Error(), nil)
		return
	}
	resp.Response(http.StatusOK, 0, "", data)
}

func CredentialAckInfo(c *gin.Context) {
	resp := Gin{C: c}
	req := &message.CredentialACK{}
	err := c.Bind(req)
	if err != nil {
		middleware.Log.Errorf("CredentialAck err:%s", err)
		resp.Response(http.StatusOK, 0, err.Error(), nil)
		return
	}
	data, err := SendMsg(message.CredentialACKType, req)
	if err != nil {
		middleware.Log.Errorf("CredentialAck msg type:%d,err:%s", message.CredentialACKType, err)
		resp.Response(http.StatusOK, 0, err.Error(), nil)
		return
	}
	resp.Response(http.StatusOK, 0, "", data)
}

func RequestProof(c *gin.Context) {
	resp := Gin{C: c}
	req := &message.RequestCredential{}
	err := c.Bind(req)
	if err != nil {
		middleware.Log.Errorf("RequestCredential err:%s", err)
		resp.Response(http.StatusOK, 0, err.Error(), nil)
		return
	}
	data, err := SendMsg(message.RequestPresentationType, req)
	if err != nil {
		middleware.Log.Errorf("RequestCredential msg type:%d,err:%s", message.RequestCredentialType, err)
		resp.Response(http.StatusOK, 0, err.Error(), nil)
		return
	}
	resp.Response(http.StatusOK, 0, "", data)
}

func PresentProof(c *gin.Context) {
	resp := Gin{C: c}
	req := &message.Presentation{}
	err := c.Bind(req)
	if err != nil {
		middleware.Log.Errorf("Presentation err:%s", err)
		resp.Response(http.StatusOK, 0, err.Error(), nil)
		return
	}
	data, err := SendMsg(message.PresentationType, req)
	if err != nil {
		middleware.Log.Errorf("Presentation msg type:%d,err:%s", message.PresentationType, err)
		resp.Response(http.StatusOK, 0, err.Error(), nil)
		return
	}
	resp.Response(http.StatusOK, 0, "", data)
}

func PresentationACKInfo(c *gin.Context) {
	resp := Gin{C: c}
	req := &message.GeneralACK{}
	err := c.Bind(req)
	if err != nil {
		middleware.Log.Errorf("PresentationACK err:%s", err)
		resp.Response(http.StatusOK, 0, err.Error(), nil)
		return
	}
	data, err := SendMsg(message.PresentationACKType, req)
	if err != nil {
		middleware.Log.Errorf("PresentationACK msg type:%d,err:%s", message.PresentationACKType, err)
		resp.Response(http.StatusOK, 0, err.Error(), nil)
		return
	}
	resp.Response(http.StatusOK, 0, "", data)
}

func SendGeneralMsg(c *gin.Context) {
	resp := Gin{C: c}
	req := &message.BasicMessage{}
	err := c.Bind(req)
	if err != nil {
		middleware.Log.Errorf("SendGeneralMsg err:%s", err)
		resp.Response(http.StatusOK, 0, err.Error(), nil)
		return
	}
	data, err := SendMsg(message.GeneralMsgType, req)
	if err != nil {
		middleware.Log.Errorf("SendGeneralMsg msg type:%d,err:%s", message.GeneralMsgType, err)
		resp.Response(http.StatusOK, 0, err.Error(), nil)
		return
	}
	resp.Response(http.StatusOK, 0, "", data)
}
