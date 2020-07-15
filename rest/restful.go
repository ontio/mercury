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
	resolveRequest(c, message.InvitationType)
	//resp := Gin{C: c}
	//req := &message.Invitation{}
	//var err error
	//var ok bool
	//if EnablePackage {
	//	msg, err := ParseMsg(c)
	//	if err != nil {
	//		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
	//		return
	//	}
	//	req, ok = msg.(*message.Invitation)
	//	if !ok {
	//		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, "msg parse error", nil)
	//		return
	//	}
	//} else {
	//	err = c.Bind(req)
	//}
	//if err != nil {
	//	middleware.Log.Errorf("Invite err:%s", err)
	//	resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
	//	return
	//}
	//data, err := SendMsg(message.InvitationType, req)
	//if err != nil {
	//	middleware.Log.Errorf("Invite err:%s", err)
	//	resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
	//	return
	//}
	//resp.Response(http.StatusOK, message.SUCCEED_CODE, "", data)
}

func resolveRequest(c *gin.Context, messageType message.MessageType) {
	resp := Gin{C: c}
	req, err := getReqByMessageType(int(messageType))
	if err != nil {
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
		return
	}
	var ok bool
	if EnablePackage {
		msg, err := ParseMsg(c)
		if err != nil {
			resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
			return
		}
		//req, ok = msg.(*message.Invitation)
		switch messageType {
		case message.InvitationType:
			req = msg.(*message.Invitation)

		case message.ConnectionRequestType:
			req = msg.(*message.ConnectionRequest)

		case message.ConnectionResponseType:
			req = msg.(*message.ConnectionResponse)

		case message.ConnectionACKType:
			req = msg.(*message.ConnectionACK)

		case message.DisconnectType, message.SendDisconnectType:
			req = msg.(*message.DisconnectRequest)

		case message.SendProposalCredentialType:
			req = msg.(*message.ProposalCredential)

		case message.OfferCredentialType:
			req = msg.(*message.OfferCredential)

		case message.ProposalCredentialType:
			req = msg.(*message.ProposalCredential)

		case message.SendRequestCredentialType:
			req = msg.(*message.RequestCredential)

		case message.RequestCredentialType:
			req = msg.(*message.RequestCredential)

		case message.IssueCredentialType:
			req = msg.(*message.IssueCredential)

		case message.RequestPresentationType:
			req = msg.(*message.RequestPresentation)

		case message.SendRequestPresentationType:
			req = msg.(*message.RequestPresentation)

		case message.SendGeneralMsgType:
			req = msg.(*message.BasicMessage)

		case message.QueryGeneralMessageType:
			req = msg.(*message.QueryGeneralMessageRequest)

		case message.QueryCredentialType:
			req = msg.(*message.QueryCredentialRequest)

		case message.QueryPresentationType:
			req = msg.(*message.QueryPresentationRequest)

		default:
			resp.Response(http.StatusOK, message.ERROR_CODE_INNER, "not a supported message type", nil)
			return
		}

		if !ok {
			resp.Response(http.StatusOK, message.ERROR_CODE_INNER, "msg parse error", nil)
			return
		}
	} else {
		err = c.Bind(req)
	}
	if err != nil {
		middleware.Log.Errorf("Bind err:%s", err)
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
		return
	}
	data, err := SendMsg(messageType, req)
	if err != nil {
		middleware.Log.Errorf("SendMsg err:%s", err)
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
		return
	}
	resp.Response(http.StatusOK, message.SUCCEED_CODE, "", data)
}

func ConnectRequest(c *gin.Context) {
	resolveRequest(c, message.ConnectionRequestType)

	//resp := Gin{C: c}
	//req := &message.ConnectionRequest{}
	//var err error
	//var ok bool
	//if EnablePackage {
	//	msg, err := ParseMsg(c)
	//	if err != nil {
	//		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
	//		return
	//	}
	//	req, ok = msg.(*message.ConnectionRequest)
	//	if !ok {
	//		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, "msg parse error", nil)
	//		return
	//	}
	//} else {
	//	err = c.Bind(req)
	//}
	//if err != nil {
	//	middleware.Log.Errorf("ConnectionRequest err:%s", err)
	//	resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
	//	return
	//}
	//data, err := SendMsg(message.ConnectionRequestType, req)
	//if err != nil {
	//	middleware.Log.Errorf("connect err:%s", err.Error())
	//	resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
	//	return
	//}
	//resp.Response(http.StatusOK, message.SUCCEED_CODE, "", data)
}

func ConnectResponse(c *gin.Context) {
	resolveRequest(c, message.ConnectionResponseType)

	//resp := Gin{C: c}
	//req := &message.ConnectionResponse{}
	//var err error
	//var ok bool
	//if EnablePackage {
	//	msg, err := ParseMsg(c)
	//	if err != nil {
	//		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
	//		return
	//	}
	//	req, ok = msg.(*message.ConnectionResponse)
	//	if !ok {
	//		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, "msg parse error", nil)
	//		return
	//	}
	//} else {
	//	err = c.Bind(req)
	//}
	//if err != nil {
	//	middleware.Log.Errorf("ConnectResponse err:%s", err)
	//	resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
	//	return
	//}
	//data, err := SendMsg(message.ConnectionResponseType, req)
	//if err != nil {
	//	middleware.Log.Errorf("connect err:%s", err.Error())
	//	resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
	//	return
	//}
	//resp.Response(http.StatusOK, message.SUCCEED_CODE, "", data)
}

func ConnectAck(c *gin.Context) {
	resolveRequest(c, message.ConnectionACKType)
	//resp := Gin{C: c}
	//req := &message.ConnectionACK{}
	//var err error
	//var ok bool
	//if EnablePackage {
	//	msg, err := ParseMsg(c)
	//	if err != nil {
	//		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
	//		return
	//	}
	//	req, ok = msg.(*message.ConnectionACK)
	//	if !ok {
	//		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, "msg parse error", nil)
	//		return
	//	}
	//} else {
	//	err = c.Bind(req)
	//}
	//if err != nil {
	//	middleware.Log.Errorf("ConnectAck err:%s", err)
	//	resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
	//	return
	//}
	//data, err := SendMsg(message.ConnectionACKType, req)
	//if err != nil {
	//	middleware.Log.Errorf("ConnectAck msg type:%d,err:%s", message.ConnectionACKType, err)
	//	resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
	//	return
	//}
	//resp.Response(http.StatusOK, message.SUCCEED_CODE, "", data)
}

func Disconnect(c *gin.Context) {
	resolveRequest(c, message.DisconnectType)
	//
	//resp := Gin{C: c}
	//req := &message.DisconnectRequest{}
	//var err error
	//var ok bool
	//if EnablePackage {
	//	msg, err := ParseMsg(c)
	//	if err != nil {
	//		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
	//		return
	//	}
	//	req, ok = msg.(*message.DisconnectRequest)
	//	if !ok {
	//		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, "msg parse error", nil)
	//		return
	//	}
	//} else {
	//	err = c.Bind(req)
	//}
	//if err != nil {
	//	middleware.Log.Errorf("Disconnect err:%s", err)
	//	resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
	//	return
	//}
	//data, err := SendMsg(message.DisconnectType, req)
	//if err != nil {
	//	middleware.Log.Errorf("ProposalCredentialReq msg type:%d,err:%s", message.DisconnectType, err)
	//	resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
	//	return
	//}
	//resp.Response(http.StatusOK, message.SUCCEED_CODE, "", data)
}

func SendDisconnect(c *gin.Context) {
	resolveRequest(c, message.SendDisconnectType)
	//
	//resp := Gin{C: c}
	//req := &message.DisconnectRequest{}
	//var err error
	//var ok bool
	//if EnablePackage {
	//	msg, err := ParseMsg(c)
	//	if err != nil {
	//		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
	//		return
	//	}
	//	req, ok = msg.(*message.DisconnectRequest)
	//	if !ok {
	//		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, "msg parse error", nil)
	//		return
	//	}
	//} else {
	//	err = c.Bind(req)
	//}
	//if err != nil {
	//	middleware.Log.Errorf("Disconnect err:%s", err)
	//	resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
	//	return
	//}
	//data, err := SendMsg(message.SendDisconnectType, req)
	//if err != nil {
	//	middleware.Log.Errorf("ProposalCredentialReq msg type:%d,err:%s", message.SendDisconnectType, err)
	//	resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
	//	return
	//}
	//resp.Response(http.StatusOK, message.SUCCEED_CODE, "", data)
}

func SendProposalCredentialReq(c *gin.Context) {
	resolveRequest(c, message.SendProposalCredentialType)
	//
	//resp := Gin{C: c}
	//req := &message.ProposalCredential{}
	//var err error
	//var ok bool
	//if EnablePackage {
	//	msg, err := ParseMsg(c)
	//	if err != nil {
	//		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
	//		return
	//	}
	//	req, ok = msg.(*message.ProposalCredential)
	//	if !ok {
	//		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, "msg parse error", nil)
	//		return
	//	}
	//} else {
	//	err = c.Bind(req)
	//}
	//if err != nil {
	//	middleware.Log.Errorf("ProposalCredential err:%s", err)
	//	resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
	//	return
	//}
	//data, err := SendMsg(message.SendProposalCredentialType, req)
	//if err != nil {
	//	middleware.Log.Errorf("ProposalCredentialReq msg type:%d,err:%s", message.SendProposalCredentialType, err)
	//	resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
	//	return
	//}
	//resp.Response(http.StatusOK, message.SUCCEED_CODE, "", data)
}
func OfferCredential(c *gin.Context) {
	resolveRequest(c, message.OfferCredentialType)

	//resp := Gin{C: c}
	//req := &message.OfferCredential{}
	//var err error
	//var ok bool
	//if EnablePackage {
	//	msg, err := ParseMsg(c)
	//	if err != nil {
	//		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
	//		return
	//	}
	//	req, ok = msg.(*message.OfferCredential)
	//	if !ok {
	//		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, "msg parse error", nil)
	//		return
	//	}
	//} else {
	//	err = c.Bind(req)
	//}
	//if err != nil {
	//	middleware.Log.Errorf("OfferCredential err:%s", err)
	//	resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
	//	return
	//}
	//data, err := SendMsg(message.OfferCredentialType, req)
	//if err != nil {
	//	middleware.Log.Errorf("ProposalCredentialReq msg type:%d,err:%s", message.OfferCredentialType, err)
	//	resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
	//	return
	//}
	//resp.Response(http.StatusOK, message.SUCCEED_CODE, "", data)
}

func ProposalCredentialReq(c *gin.Context) {
	resolveRequest(c, message.ProposalCredentialType)
	//
	//resp := Gin{C: c}
	//req := &message.ProposalCredential{}
	//var err error
	//var ok bool
	//if EnablePackage {
	//	msg, err := ParseMsg(c)
	//	if err != nil {
	//		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
	//		return
	//	}
	//	req, ok = msg.(*message.ProposalCredential)
	//	if !ok {
	//		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, "msg parse error", nil)
	//		return
	//	}
	//} else {
	//	err = c.Bind(req)
	//}
	//if err != nil {
	//	middleware.Log.Errorf("ProposalCredential err:%s", err)
	//	resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
	//	return
	//}
	//data, err := SendMsg(message.ProposalCredentialType, req)
	//if err != nil {
	//	middleware.Log.Errorf("ProposalCredentialReq msg type:%d,err:%s", message.ProposalCredentialType, err)
	//	resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
	//	return
	//}
	//resp.Response(http.StatusOK, message.SUCCEED_CODE, "", data)
}

func SendRequestCredential(c *gin.Context) {
	resolveRequest(c, message.SendRequestCredentialType)

	//resp := Gin{C: c}
	//req := &message.RequestCredential{}
	//err := c.Bind(req)
	//if err != nil {
	//	middleware.Log.Errorf("SendRequestCredential err:%s", err)
	//	resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
	//	return
	//}
	//data, err := SendMsg(message.SendRequestCredentialType, req)
	//if err != nil {
	//	middleware.Log.Errorf("SendRequestCredential msg type:%d,err:%s", message.SendRequestCredentialType, err)
	//	resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
	//	return
	//}
	//resp.Response(http.StatusOK, message.SUCCEED_CODE, "", data)
}
func RequestCredential(c *gin.Context) {
	resolveRequest(c, message.RequestCredentialType)
	//
	//resp := Gin{C: c}
	//req := &message.RequestCredential{}
	//var err error
	//var ok bool
	//if EnablePackage {
	//	msg, err := ParseMsg(c)
	//	if err != nil {
	//		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
	//		return
	//	}
	//	req, ok = msg.(*message.RequestCredential)
	//	if !ok {
	//		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, "msg parse error", nil)
	//		return
	//	}
	//} else {
	//	err = c.Bind(req)
	//}
	//if err != nil {
	//	middleware.Log.Errorf("ProposalCredential err:%s", err)
	//	resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
	//	return
	//}
	//data, err := SendMsg(message.RequestCredentialType, req)
	//if err != nil {
	//	middleware.Log.Errorf("ProposalCredential msg type:%d,err:%s", message.RequestCredentialType, err)
	//	resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
	//	return
	//}
	//resp.Response(http.StatusOK, message.ERROR_CODE_INNER, "", data)
}

func IssueCredential(c *gin.Context) {
	resolveRequest(c, message.IssueCredentialType)

	//resp := Gin{C: c}
	//req := &message.IssueCredential{}
	//var err error
	//var ok bool
	//if EnablePackage {
	//	msg, err := ParseMsg(c)
	//	if err != nil {
	//		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
	//		return
	//	}
	//	req, ok = msg.(*message.IssueCredential)
	//	if !ok {
	//		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, "msg parse error", nil)
	//		return
	//	}
	//} else {
	//	err = c.Bind(req)
	//}
	//if err != nil {
	//	middleware.Log.Errorf("OfferCredential err:%s", err)
	//	resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
	//	return
	//}
	//data, err := SendMsg(message.IssueCredentialType, req)
	//if err != nil {
	//	middleware.Log.Errorf("OfferCredential msg type:%d,err:%s", message.IssueCredentialType, err)
	//	resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
	//	return
	//}
	//resp.Response(http.StatusOK, message.SUCCEED_CODE, "", data)
}

func CredentialAckInfo(c *gin.Context) {
	resolveRequest(c, message.CredentialACKType)

	//resp := Gin{C: c}
	//req := &message.CredentialACK{}
	//err := c.Bind(req)
	//if err != nil {
	//	middleware.Log.Errorf("CredentialAck err:%s", err)
	//	resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
	//	return
	//}
	//data, err := SendMsg(message.CredentialACKType, req)
	//if err != nil {
	//	middleware.Log.Errorf("CredentialAck msg type:%d,err:%s", message.CredentialACKType, err)
	//	resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
	//	return
	//}
	//resp.Response(http.StatusOK, message.SUCCEED_CODE, "", data)
}

func RequestPresentation(c *gin.Context) {
	resolveRequest(c, message.RequestPresentationType)

	//resp := Gin{C: c}
	//req := &message.RequestPresentation{}
	//var err error
	//var ok bool
	//if EnablePackage {
	//	msg, err := ParseMsg(c)
	//	if err != nil {
	//		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
	//		return
	//	}
	//	req, ok = msg.(*message.RequestPresentation)
	//	if !ok {
	//		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, "msg parse error", nil)
	//		return
	//	}
	//} else {
	//	err = c.Bind(req)
	//}
	//if err != nil {
	//	middleware.Log.Errorf("RequestCredential err:%s", err)
	//	resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
	//	return
	//}
	//data, err := SendMsg(message.RequestPresentationType, req)
	//if err != nil {
	//	middleware.Log.Errorf("RequestCredential msg type:%d,err:%s", message.RequestCredentialType, err)
	//	resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
	//	return
	//}
	//resp.Response(http.StatusOK, message.SUCCEED_CODE, "", data)
}

func SendRequestPresentation(c *gin.Context) {
	resolveRequest(c, message.SendRequestPresentationType)

	//resp := Gin{C: c}
	//req := &message.RequestPresentation{}
	//var err error
	//var ok bool
	//if EnablePackage {
	//	msg, err := ParseMsg(c)
	//	if err != nil {
	//		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
	//		return
	//	}
	//	req, ok = msg.(*message.RequestPresentation)
	//	if !ok {
	//		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, "msg parse error", nil)
	//		return
	//	}
	//} else {
	//	err = c.Bind(req)
	//}
	//if err != nil {
	//	middleware.Log.Errorf("RequestPresentationType err:%s", err)
	//	resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
	//	return
	//}
	//data, err := SendMsg(message.SendRequestPresentationType, req)
	//if err != nil {
	//	middleware.Log.Errorf("RequestCredential msg type:%d,err:%s", message.SendRequestPresentationType, err)
	//	resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
	//	return
	//}
	//resp.Response(http.StatusOK, message.SUCCEED_CODE, "", data)
}

func Presentation(c *gin.Context) {
	resolveRequest(c, message.PresentationType)
	//
	//resp := Gin{C: c}
	//req := &message.Presentation{}
	//err := c.Bind(req)
	//if err != nil {
	//	middleware.Log.Errorf("Presentation err:%s", err)
	//	resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
	//	return
	//}
	//data, err := SendMsg(message.PresentationType, req)
	//if err != nil {
	//	middleware.Log.Errorf("Presentation msg type:%d,err:%s", message.PresentationType, err)
	//	resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
	//	return
	//}
	//resp.Response(http.StatusOK, message.SUCCEED_CODE, "", data)
}

func PresentationAckInfo(c *gin.Context) {
	resolveRequest(c, message.PresentationACKType)

	//resp := Gin{C: c}
	//req := &message.PresentationACK{}
	//err := c.Bind(req)
	//if err != nil {
	//	middleware.Log.Errorf("PresentationACK err:%s", err)
	//	resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
	//	return
	//}
	//data, err := SendMsg(message.PresentationACKType, req)
	//if err != nil {
	//	middleware.Log.Errorf("PresentationACK msg type:%d,err:%s", message.PresentationACKType, err)
	//	resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
	//	return
	//}
	//resp.Response(http.StatusOK, message.SUCCEED_CODE, "", data)
}

func SendGeneralMsg(c *gin.Context) {
	resolveRequest(c, message.SendGeneralMsgType)

	//resp := Gin{C: c}
	//req := &message.BasicMessage{}
	//var err error
	//var ok bool
	//if EnablePackage {
	//	msg, err := ParseMsg(c)
	//	if err != nil {
	//		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
	//		return
	//	}
	//	req, ok = msg.(*message.BasicMessage)
	//	if !ok {
	//		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, "msg parse error", nil)
	//		return
	//	}
	//} else {
	//	err = c.Bind(req)
	//}
	//if err != nil {
	//	middleware.Log.Errorf("SendGeneralMsg err:%s", err)
	//	resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
	//	return
	//}
	//data, err := SendMsg(message.SendGeneralMsgType, req)
	//if err != nil {
	//	middleware.Log.Errorf("SendGeneralMsg msg type:%d,err:%s", message.SendGeneralMsgType, err)
	//	resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
	//	return
	//}
	//resp.Response(http.StatusOK, message.SUCCEED_CODE, "", data)
}

func ReceiveGeneralMsg(c *gin.Context) {
	resolveRequest(c, message.ReceiveGeneralMsgType)

	//resp := Gin{C: c}
	//req := &message.BasicMessage{}
	//err := c.Bind(req)
	//if err != nil {
	//	middleware.Log.Errorf("ReceiveGeneralMsg err:%s", err)
	//	resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
	//	return
	//}
	//data, err := SendMsg(message.ReceiveGeneralMsgType, req)
	//if err != nil {
	//	middleware.Log.Errorf("SendGeneralMsg msg type:%d,err:%s", message.ReceiveGeneralMsgType, err)
	//	resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
	//	return
	//}
	//resp.Response(http.StatusOK, message.SUCCEED_CODE, "", data)
}
func QueryGeneralMsg(c *gin.Context) {
	resolveRequest(c, message.QueryGeneralMessageType)

	//resp := Gin{C: c}
	//req := &message.QueryGeneralMessageRequest{}
	//var err error
	//var ok bool
	//if EnablePackage {
	//	msg, err := ParseMsg(c)
	//	if err != nil {
	//		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
	//		return
	//	}
	//	req, ok = msg.(*message.QueryGeneralMessageRequest)
	//	if !ok {
	//		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, "msg parse error", nil)
	//		return
	//	}
	//} else {
	//	err = c.Bind(req)
	//}
	//if err != nil {
	//	middleware.Log.Errorf("QueryGeneralMessage err:%s", err)
	//	resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
	//	return
	//}
	//data, err := SendMsg(message.QueryGeneralMessageType, req)
	//if err != nil {
	//	middleware.Log.Errorf("SendGeneralMsg msg type:%d,err:%s", message.QueryGeneralMessageType, err)
	//	resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
	//	return
	//}
	//resp.Response(http.StatusOK, message.SUCCEED_CODE, "", data)
}

func QueryCredential(c *gin.Context) {
	resolveRequest(c, message.QueryCredentialType)
	//
	//resp := Gin{C: c}
	//
	//req := &message.QueryCredentialRequest{}
	//var err error
	//var ok bool
	//if EnablePackage {
	//	msg, err := ParseMsg(c)
	//	if err != nil {
	//		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
	//		return
	//	}
	//	req, ok = msg.(*message.QueryCredentialRequest)
	//	if !ok {
	//		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, "msg parse error", nil)
	//		return
	//	}
	//} else {
	//	err = c.Bind(req)
	//}
	//if err != nil {
	//	middleware.Log.Errorf("QueryCredential err:%s", err)
	//	resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
	//	return
	//}
	//data, err := SendMsg(message.QueryCredentialType, req)
	//if err != nil {
	//	middleware.Log.Errorf("QueryCredential msg type:%d,err:%s", message.QueryCredentialType, err)
	//	resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
	//	return
	//}
	//resp.Response(http.StatusOK, message.SUCCEED_CODE, "", data)
}

func QueryPresentation(c *gin.Context) {
	resolveRequest(c, message.QueryPresentationType)

	//resp := Gin{C: c}
	//req := &message.QueryPresentationRequest{}
	//var err error
	//var ok bool
	//if EnablePackage {
	//	msg, err := ParseMsg(c)
	//	if err != nil {
	//		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
	//		return
	//	}
	//	req, ok = msg.(*message.QueryPresentationRequest)
	//	if !ok {
	//		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, "msg parse error", nil)
	//		return
	//	}
	//} else {
	//	err = c.Bind(req)
	//}
	//if err != nil {
	//	middleware.Log.Errorf("QueryPresentation err:%s", err)
	//	resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
	//	return
	//}
	//data, err := SendMsg(message.QueryPresentationType, req)
	//if err != nil {
	//	middleware.Log.Errorf("QueryPresentation msg type:%d,err:%s", message.QueryPresentationType, err)
	//	resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
	//	return
	//}
	//resp.Response(http.StatusOK, message.SUCCEED_CODE, "", data)
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
	/*var req interface{}
	switch msg.Message.MsgType {
	case int(message.InvitationType):
		req = &message.Invitation{}

	case int(message.ConnectionRequestType):
		req = &message.ConnectionRequest{}

	case int(message.ConnectionResponseType):
		req = &message.ConnectionResponse{}

	case int(message.ConnectionACKType):
		req = &message.ConnectionACK{}

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
	}*/
	req, err := getReqByMessageType(msg.Message.MsgType)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(msg.Message.Data, req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func getReqByMessageType(t int) (interface{}, error) {
	var req interface{}
	switch t {
	case int(message.InvitationType):
		req = &message.Invitation{}

	case int(message.ConnectionRequestType):
		req = &message.ConnectionRequest{}

	case int(message.ConnectionResponseType):
		req = &message.ConnectionResponse{}

	case int(message.ConnectionACKType):
		req = &message.ConnectionACK{}
	case int(message.DisconnectType), int(message.SendDisconnectType):
		req = &message.DisconnectRequest{}

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
		return nil, fmt.Errorf("msg type err:%d", t)
	}
	return req, nil
}
