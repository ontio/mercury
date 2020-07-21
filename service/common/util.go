package common

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"git.ont.io/ontid/otf/common/message"
	"git.ont.io/ontid/otf/common/packager/ecdsa"
	"github.com/gin-gonic/gin"
)

var (
	EnablePackage bool
)

func ReverseConnection(conn message.Connection) message.Connection {
	return message.Connection{
		MyDid:       conn.TheirDid,
		MyRouter:    conn.TheirRouter,
		TheirDid:    conn.MyDid,
		TheirRouter: conn.MyRouter,
	}
}

func ParseRouterMsg(c *gin.Context, packager *ecdsa.Packager) ([]byte, error) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		return nil, err
	}
	msg, err := packager.UnpackMessage(body)
	if err != nil {
		return nil, err
	}
	if msg.Message == nil {
		return nil, fmt.Errorf("msg is nil")
	}
	return msg.Message.Data, nil
}

func ParseMessage(enablePackage bool, ctx *gin.Context, packager *ecdsa.Packager, messageType message.MessageType) (interface{}, error) {
	msgObject, err := getMsgObjectByType(messageType)
	if err != nil {
		return nil, err
	}
	if enablePackage {
		data, err := ParseRouterMsg(ctx, packager)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(data, msgObject)
		if err != nil {
			return nil, err
		}
	} else {
		err = ctx.Bind(msgObject)
		if err != nil {
			return nil, err
		}
	}
	return msgObject, nil
}

func getMsgObjectByType(messageType message.MessageType) (interface{}, error) {
	var req interface{}
	switch messageType {
	case message.InvitationType:
		req = &message.Invitation{}
	case message.ConnectionRequestType:
		req = &message.ConnectionRequest{}
	case message.ConnectionResponseType:
		req = &message.ConnectionResponse{}
	case message.ConnectionACKType:
		req = &message.ConnectionACK{}
	case message.DisconnectType, message.SendDisconnectType:
		req = &message.DisconnectRequest{}
	case message.SendProposalCredentialType:
		req = &message.ProposalCredential{}
	case message.OfferCredentialType:
		req = &message.OfferCredential{}
	case message.ProposalCredentialType:
		req = &message.ProposalCredential{}
	case message.SendRequestCredentialType:
		req = &message.RequestCredential{}
	case message.RequestCredentialType:
		req = &message.RequestCredential{}
	case message.IssueCredentialType:
		req = &message.IssueCredential{}
	case message.CredentialACKType:
		req = &message.CredentialACK{}
	case message.RequestPresentationType:
		req = &message.RequestPresentation{}
	case message.SendRequestPresentationType:
		req = &message.RequestPresentation{}
	case message.PresentationType:
		req = &message.Presentation{}
	case message.PresentationACKType:
		req = &message.PresentationACK{}
	case message.SendGeneralMsgType:
		req = &message.BasicMessage{}
	case message.QueryGeneralMessageType:
		req = &message.QueryGeneralMessageRequest{}
	case message.QueryCredentialType:
		req = &message.QueryCredentialRequest{}
	case message.QueryPresentationType:
		req = &message.QueryPresentationRequest{}
	default:
		return nil, fmt.Errorf("msg type err:%v", messageType)
	}
	return req, nil
}
