package rest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"git.ont.io/ontid/otf/common/log"
	"git.ont.io/ontid/otf/common/message"
	"github.com/gin-gonic/gin"
)

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

		case message.CredentialACKType:
			req = msg.(*message.CredentialACK)

		case message.RequestPresentationType:
			req = msg.(*message.RequestPresentation)

		case message.SendRequestPresentationType:
			req = msg.(*message.RequestPresentation)

		case message.PresentationType:
			req = msg.(*message.Presentation)

		case message.PresentationACKType:
			req = msg.(*message.PresentationACK)

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
		log.Errorf("Bind err:%s", err)
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
		return
	}
	data, err := SendMsg(messageType, req)
	if err != nil {
		log.Errorf("SendMsg err:%s", err)
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
		return
	}
	resp.Response(http.StatusOK, message.SUCCEED_CODE, "", data)
}

func Invite(c *gin.Context) {
	resolveRequest(c, message.InvitationType)
}

func ConnectRequest(c *gin.Context) {
	resolveRequest(c, message.ConnectionRequestType)
}

func ConnectResponse(c *gin.Context) {
	resolveRequest(c, message.ConnectionResponseType)
}

func ConnectAck(c *gin.Context) {
	resolveRequest(c, message.ConnectionACKType)
}

func Disconnect(c *gin.Context) {
	resolveRequest(c, message.DisconnectType)
}

func SendDisconnect(c *gin.Context) {
	resolveRequest(c, message.SendDisconnectType)
}

func SendProposalCredentialReq(c *gin.Context) {
	resolveRequest(c, message.SendProposalCredentialType)
}

func OfferCredential(c *gin.Context) {
	resolveRequest(c, message.OfferCredentialType)
}

func ProposalCredentialReq(c *gin.Context) {
	resolveRequest(c, message.ProposalCredentialType)
}

func SendRequestCredential(c *gin.Context) {
	resolveRequest(c, message.SendRequestCredentialType)
}

func RequestCredential(c *gin.Context) {
	resolveRequest(c, message.RequestCredentialType)
}

func IssueCredential(c *gin.Context) {
	resolveRequest(c, message.IssueCredentialType)
}

func CredentialAckInfo(c *gin.Context) {
	resolveRequest(c, message.CredentialACKType)
}

func RequestPresentation(c *gin.Context) {
	resolveRequest(c, message.RequestPresentationType)
}

func SendRequestPresentation(c *gin.Context) {
	resolveRequest(c, message.SendRequestPresentationType)
}

func Presentation(c *gin.Context) {
	resolveRequest(c, message.PresentationType)
}

func PresentationAckInfo(c *gin.Context) {
	resolveRequest(c, message.PresentationACKType)
}

func SendGeneralMsg(c *gin.Context) {
	resolveRequest(c, message.SendGeneralMsgType)
}

func ReceiveGeneralMsg(c *gin.Context) {
	resolveRequest(c, message.ReceiveGeneralMsgType)
}
func QueryGeneralMsg(c *gin.Context) {
	resolveRequest(c, message.QueryGeneralMessageType)
}

func QueryCredential(c *gin.Context) {
	resolveRequest(c, message.QueryCredentialType)
}

func QueryPresentation(c *gin.Context) {
	resolveRequest(c, message.QueryPresentationType)
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

	case int(message.CredentialACKType):
		req = &message.CredentialACK{}

	case int(message.RequestPresentationType):
		req = &message.RequestPresentation{}

	case int(message.SendRequestPresentationType):
		req = &message.RequestPresentation{}

	case int(message.PresentationType):
		req = &message.Presentation{}

	case int(message.PresentationACKType):
		req = &message.PresentationACK{}

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
