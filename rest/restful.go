package rest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"git.ont.io/ontid/otf/message"
	"git.ont.io/ontid/otf/middleware"
	"github.com/gin-gonic/gin"
)

func Invite(c *gin.Context) {
	resp := Gin{C: c}
	req := &message.Invitation{}
	var err error
	var ok bool
	if EnablePackage {
		msg, err := ParseMsg(c)
		if err != nil {
			resp.Response(http.StatusOK, 0, err.Error(), nil)
			return
		}
		req, ok = msg.(*message.Invitation)
		if !ok {
			resp.Response(http.StatusOK, 0, "msg parse error", nil)
			return
		}
	} else {
		err = c.Bind(req)
	}
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

func ConnectRequest(c *gin.Context) {
	resp := Gin{C: c}
	req := &message.ConnectionRequest{}
	var err error
	var ok bool
	if EnablePackage {
		msg, err := ParseMsg(c)
		if err != nil {
			resp.Response(http.StatusOK, 0, err.Error(), nil)
			return
		}
		req, ok = msg.(*message.ConnectionRequest)
		if !ok {
			resp.Response(http.StatusOK, 0, "msg parse error", nil)
			return
		}
	} else {
		err = c.Bind(req)
	}
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
	var err error
	var ok bool
	if EnablePackage {
		msg, err := ParseMsg(c)
		if err != nil {
			resp.Response(http.StatusOK, 0, err.Error(), nil)
			return
		}
		req, ok = msg.(*message.ProposalCredential)
		if !ok {
			resp.Response(http.StatusOK, 0, "msg parse error", nil)
			return
		}
	} else {
		err = c.Bind(req)
	}
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
	var err error
	var ok bool
	if EnablePackage {
		msg, err := ParseMsg(c)
		if err != nil {
			resp.Response(http.StatusOK, 0, err.Error(), nil)
			return
		}
		req, ok = msg.(*message.OfferCredential)
		if !ok {
			resp.Response(http.StatusOK, 0, "msg parse error", nil)
			return
		}
	} else {
		err = c.Bind(req)
	}
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
	var err error
	var ok bool
	if EnablePackage {
		msg, err := ParseMsg(c)
		if err != nil {
			resp.Response(http.StatusOK, 0, err.Error(), nil)
			return
		}
		req, ok = msg.(*message.ProposalCredential)
		if !ok {
			resp.Response(http.StatusOK, 0, "msg parse error", nil)
			return
		}
	} else {
		err = c.Bind(req)
	}
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
	var err error
	var ok bool
	if EnablePackage {
		msg, err := ParseMsg(c)
		if err != nil {
			resp.Response(http.StatusOK, 0, err.Error(), nil)
			return
		}
		req, ok = msg.(*message.RequestCredential)
		if !ok {
			resp.Response(http.StatusOK, 0, "msg parse error", nil)
			return
		}
	} else {
		err = c.Bind(req)
	}
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
	var err error
	var ok bool
	if EnablePackage {
		msg, err := ParseMsg(c)
		if err != nil {
			resp.Response(http.StatusOK, 0, err.Error(), nil)
			return
		}
		req, ok = msg.(*message.IssueCredential)
		if !ok {
			resp.Response(http.StatusOK, 0, "msg parse error", nil)
			return
		}
	} else {
		err = c.Bind(req)
	}
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
	var err error
	var ok bool
	if EnablePackage {
		msg, err := ParseMsg(c)
		if err != nil {
			resp.Response(http.StatusOK, 0, err.Error(), nil)
			return
		}
		req, ok = msg.(*message.RequestPresentation)
		if !ok {
			resp.Response(http.StatusOK, 0, "msg parse error", nil)
			return
		}
	} else {
		err = c.Bind(req)
	}
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
	var err error
	var ok bool
	if EnablePackage {
		msg, err := ParseMsg(c)
		if err != nil {
			resp.Response(http.StatusOK, 0, err.Error(), nil)
			return
		}
		req, ok = msg.(*message.RequestPresentation)
		if !ok {
			resp.Response(http.StatusOK, 0, "msg parse error", nil)
			return
		}
	} else {
		err = c.Bind(req)
	}
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
	var err error
	var ok bool
	if EnablePackage {
		msg, err := ParseMsg(c)
		if err != nil {
			resp.Response(http.StatusOK, 0, err.Error(), nil)
			return
		}
		req, ok = msg.(*message.BasicMessage)
		if !ok {
			resp.Response(http.StatusOK, 0, "msg parse error", nil)
			return
		}
	} else {
		err = c.Bind(req)
	}
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
func QueryGeneralMsg(c *gin.Context) {
	resp := Gin{C: c}
	req := &message.QueryGeneralMessageRequest{}
	var err error
	var ok bool
	if EnablePackage {
		msg, err := ParseMsg(c)
		if err != nil {
			resp.Response(http.StatusOK, 0, err.Error(), nil)
			return
		}
		req, ok = msg.(*message.QueryGeneralMessageRequest)
		if !ok {
			resp.Response(http.StatusOK, 0, "msg parse error", nil)
			return
		}
	} else {
		err = c.Bind(req)
	}
	if err != nil {
		middleware.Log.Errorf("QueryGeneralMessage err:%s", err)
		resp.Response(http.StatusOK, 0, err.Error(), nil)
		return
	}
	data, err := SendMsg(message.QueryGeneralMessageType, req)
	if err != nil {
		middleware.Log.Errorf("SendGeneralMsg msg type:%d,err:%s", message.QueryGeneralMessageType, err)
		resp.Response(http.StatusOK, 0, err.Error(), nil)
		return
	}
	resp.Response(http.StatusOK, 0, "", data)
}

func QueryCredential(c *gin.Context) {
	resp := Gin{C: c}

	req := &message.QueryCredentialRequest{}
	var err error
	var ok bool
	if EnablePackage {
		msg, err := ParseMsg(c)
		if err != nil {
			resp.Response(http.StatusOK, 0, err.Error(), nil)
			return
		}
		req, ok = msg.(*message.QueryCredentialRequest)
		if !ok {
			resp.Response(http.StatusOK, 0, "msg parse error", nil)
			return
		}
	} else {
		err = c.Bind(req)
	}
	if err != nil {
		middleware.Log.Errorf("QueryCredential err:%s", err)
		resp.Response(http.StatusOK, 0, err.Error(), nil)
		return
	}
	data, err := SendMsg(message.QueryCredentialType, req)
	if err != nil {
		middleware.Log.Errorf("QueryCredential msg type:%d,err:%s", message.QueryCredentialType, err)
		resp.Response(http.StatusOK, 0, err.Error(), nil)
		return
	}
	resp.Response(http.StatusOK, 0, "", data)
}

func QueryPresentation(c *gin.Context) {
	resp := Gin{C: c}
	req := &message.QueryPresentationRequest{}
	var err error
	var ok bool
	if EnablePackage {
		msg, err := ParseMsg(c)
		if err != nil {
			resp.Response(http.StatusOK, 0, err.Error(), nil)
			return
		}
		req, ok = msg.(*message.QueryPresentationRequest)
		if !ok {
			resp.Response(http.StatusOK, 0, "msg parse error", nil)
			return
		}
	} else {
		err = c.Bind(req)
	}
	if err != nil {
		middleware.Log.Errorf("QueryPresentation err:%s", err)
		resp.Response(http.StatusOK, 0, err.Error(), nil)
		return
	}
	data, err := SendMsg(message.QueryPresentationType, req)
	if err != nil {
		middleware.Log.Errorf("QueryPresentation msg type:%d,err:%s", message.QueryPresentationType, err)
		resp.Response(http.StatusOK, 0, err.Error(), nil)
		return
	}
	resp.Response(http.StatusOK, 0, "", data)
}

func ParseMsg(c *gin.Context) (interface{}, error) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		return nil, err
	}

	msg, err := Svr.ParseMsg(body)
	if err != nil {
		return nil, err
	}
	if msg.Message == nil {
		return nil, fmt.Errorf("msg is nil")
	}
	var req interface{}
	switch msg.Message.MsgType {
	case int(message.InvitationType):
		req = &message.Invitation{}

	case int(message.ConnectionRequestType):
		req = &message.ConnectionRequest{}

	case int(message.SendProposalCredentialType):
		req = &message.ProposalCredential{}

	case int(message.OfferCredentialType):
		req = &message.OfferCredential{}

	case int(message.ProposalCredentialType):
		req = &message.ProposalCredential{}

	case int(message.SendRequestCredentialType):
		req = &message.RequestCredential{}

	case int(message.RequestCredentialType):
		req = &message.RequestCredential{}

	case int(message.IssueCredentialType):
		req = &message.IssueCredential{}

	case int(message.RequestPresentationType):
		req = &message.RequestPresentation{}

	case int(message.SendRequestPresentationType):
		req = &message.RequestPresentation{}

	case int(message.SendGeneralMsgType):
		req = &message.BasicMessage{}

	case int(message.QueryGeneralMessageType):
		req = &message.QueryGeneralMessageRequest{}

	case int(message.QueryCredentialType):
		req = &message.QueryCredentialRequest{}

	case int(message.QueryPresentationType):
		req = &message.QueryPresentationRequest{}

	default:
		return nil, fmt.Errorf("msg type err:%d", msg.Message.MsgType)
	}
	err = json.Unmarshal(msg.Message.Data, req)
	if err != nil {
		return nil, err
	}
	return req, nil
}
