package rest

import (
	"net/http"

	"git.ont.io/ontid/otf/message"
	"git.ont.io/ontid/otf/middleware"
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
	resp.Response(http.StatusOK, 0, "", data)
}

func SendConnectionReq(c *gin.Context) {
	resp := Gin{C: c}
	req := &message.ConnectionRequest{}
	err := c.Bind(req)
	if err != nil {
		middleware.Log.Errorf("Send ConnectionRequest err:%s", err)
		resp.Response(http.StatusOK, 0, err.Error(), nil)
		return
	}
	_, err = SendMsg(message.SendConnectionRequestType, req)
	if err != nil {
		middleware.Log.Errorf("Send Connection Req err:%s", err.Error())
		resp.Response(http.StatusOK, 0, err.Error(), nil)
		return
	}
	resp.Response(http.StatusOK, 0, "", nil)
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
		middleware.Log.Errorf("connect err:%s", err.Error())
		resp.Response(http.StatusOK, 0, err.Error(), nil)
		return
	}
	resp.Response(http.StatusOK, 0, "", data)
}

func ConnectResponse(c *gin.Context) {
	resp := Gin{C: c}
	req := &message.ConnectionResponse{}
	err := c.Bind(req)
	if err != nil {
		middleware.Log.Errorf("ConnectResponse err:%s", err)
		resp.Response(http.StatusOK, 0, err.Error(), nil)
		return
	}
	data, err := SendMsg(message.ConnectionResponseType, req)
	if err != nil {
		middleware.Log.Errorf("connect err:%s", err.Error())
		resp.Response(http.StatusOK, 0, err.Error(), nil)
		return
	}
	resp.Response(http.StatusOK, 0, "", data)
}

func ConnectAck(c *gin.Context) {
	resp := Gin{C: c}
	req := &message.ConnectionACK{}
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

func SendProposalCredentialReq(c *gin.Context) {
	resp := Gin{C: c}
	req := &message.ProposalCredential{}
	err := c.Bind(req)
	if err != nil {
		middleware.Log.Errorf("ProposalCredential err:%s", err)
		resp.Response(http.StatusOK, 0, err.Error(), nil)
		return
	}
	data, err := SendMsg(message.SendProposalCredentialType, req)
	if err != nil {
		middleware.Log.Errorf("ProposalCredentialReq msg type:%d,err:%s", message.SendProposalCredentialType, err)
		resp.Response(http.StatusOK, 0, err.Error(), nil)
		return
	}
	resp.Response(http.StatusOK, 0, "", data)
}
func OfferCredential(c *gin.Context) {
	resp := Gin{C: c}
	req := &message.OfferCredential{}
	err := c.Bind(req)
	if err != nil {
		middleware.Log.Errorf("OfferCredential err:%s", err)
		resp.Response(http.StatusOK, 0, err.Error(), nil)
		return
	}
	data, err := SendMsg(message.OfferCredentialType, req)
	if err != nil {
		middleware.Log.Errorf("ProposalCredentialReq msg type:%d,err:%s", message.OfferCredentialType, err)
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
	data, err := SendMsg(message.ProposalCredentialType, req)
	if err != nil {
		middleware.Log.Errorf("ProposalCredentialReq msg type:%d,err:%s", message.ProposalCredentialType, err)
		resp.Response(http.StatusOK, 0, err.Error(), nil)
		return
	}
	resp.Response(http.StatusOK, 0, "", data)
}

func SendRequestCredential(c *gin.Context) {
	resp := Gin{C: c}
	req := &message.RequestCredential{}
	err := c.Bind(req)
	if err != nil {
		middleware.Log.Errorf("SendRequestCredential err:%s", err)
		resp.Response(http.StatusOK, 0, err.Error(), nil)
		return
	}
	data, err := SendMsg(message.SendRequestCredentialType, req)
	if err != nil {
		middleware.Log.Errorf("SendRequestCredential msg type:%d,err:%s", message.SendRequestCredentialType, err)
		resp.Response(http.StatusOK, 0, err.Error(), nil)
		return
	}
	resp.Response(http.StatusOK, 0, "", data)
}
func RequestCredential(c *gin.Context) {
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
	req := &message.IssueCredential{}
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

func RequestPresentation(c *gin.Context) {
	resp := Gin{C: c}
	req := &message.RequestPresentation{}
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

func SendRequestPresentation(c *gin.Context) {
	resp := Gin{C: c}
	req := &message.RequestPresentation{}
	err := c.Bind(req)
	if err != nil {
		middleware.Log.Errorf("RequestPresentationType err:%s", err)
		resp.Response(http.StatusOK, 0, err.Error(), nil)
		return
	}
	data, err := SendMsg(message.SendRequestPresentationType, req)
	if err != nil {
		middleware.Log.Errorf("RequestCredential msg type:%d,err:%s", message.SendRequestPresentationType, err)
		resp.Response(http.StatusOK, 0, err.Error(), nil)
		return
	}
	resp.Response(http.StatusOK, 0, "", data)
}

func Presentation(c *gin.Context) {
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

func PresentationAckInfo(c *gin.Context) {
	resp := Gin{C: c}
	req := &message.PresentationACK{}
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
	data, err := SendMsg(message.SendGeneralMsgType, req)
	if err != nil {
		middleware.Log.Errorf("SendGeneralMsg msg type:%d,err:%s", message.SendGeneralMsgType, err)
		resp.Response(http.StatusOK, 0, err.Error(), nil)
		return
	}
	resp.Response(http.StatusOK, 0, "", data)
}

func ReceiveGeneralMsg(c *gin.Context) {
	resp := Gin{C: c}
	req := &message.BasicMessage{}
	err := c.Bind(req)
	if err != nil {
		middleware.Log.Errorf("ReceiveGeneralMsg err:%s", err)
		resp.Response(http.StatusOK, 0, err.Error(), nil)
		return
	}
	data, err := SendMsg(message.ReceiveGeneralMsgType, req)
	if err != nil {
		middleware.Log.Errorf("SendGeneralMsg msg type:%d,err:%s", message.ReceiveGeneralMsgType, err)
		resp.Response(http.StatusOK, 0, err.Error(), nil)
		return
	}
	resp.Response(http.StatusOK, 0, "", data)
}
