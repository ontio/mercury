package common

import (
	"encoding/json"
	"fmt"
	"git.ont.io/ontid/otf/utils"
	"io/ioutil"
	"strings"

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

func ParseMessage(enablePackage bool, ctx *gin.Context, packager *ecdsa.Packager, messageType MessageType) (interface{}, error) {
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

func getMsgObjectByType(messageType MessageType) (interface{}, error) {
	var req interface{}
	switch messageType {
	case InvitationType:
		req = &message.Invitation{}
	case ConnectionRequestType:
		req = &message.ConnectionRequest{}
	case ConnectionResponseType:
		req = &message.ConnectionResponse{}
	case ConnectionAckType:
		req = &message.ConnectionACK{}
	case DisconnectType, SendDisconnectType:
		req = &message.DisconnectRequest{}
	case SendProposalCredentialType:
		req = &message.ProposalCredential{}
	case OfferCredentialType:
		req = &message.OfferCredential{}
	case ProposalCredentialType:
		req = &message.ProposalCredential{}
	case SendRequestCredentialType:
		req = &message.RequestCredential{}
	case RequestCredentialType:
		req = &message.RequestCredential{}
	case IssueCredentialType:
		req = &message.IssueCredential{}
	case CredentialAckType:
		req = &message.CredentialACK{}
	case RequestPresentationType:
		req = &message.RequestPresentation{}
	case SendRequestPresentationType:
		req = &message.RequestPresentation{}
	case PresentationType:
		req = &message.Presentation{}
	case PresentationAckType:
		req = &message.PresentationACK{}
	case SendBasicMsgType,ReceiveBasicMsgType:
		req = &message.BasicMessage{}
	case QueryBasicMessageType:
		req = &message.QueryGeneralMessageRequest{}
	case QueryCredentialType:
		req = &message.QueryCredentialRequest{}
	case QueryPresentationType:
		req = &message.QueryPresentationRequest{}
	default:
		return nil, fmt.Errorf("msg type err:%v", messageType)
	}
	return req, nil
}

func MergeRouter(myrouters, theirrouters []string) []string {
	return append(myrouters, reverseRouter(theirrouters)...)
}

func reverseRouter(routers []string) []string {
	ret := make([]string, 0)
	for i := len(routers) - 1; i >= 0; i-- {
		ret = append(ret, routers[i])
	}
	return ret
}

func RouterLastIndexOf(did string, routers []string) (int, error) {
	for i := len(routers) - 1; i >= 0; i-- {
		if strings.EqualFold(did, utils.CutDId(routers[i])) {
			return i, nil
		}
	}
	return -1, fmt.Errorf("cannot found did:%s in routers", did)
}

func IsReceiver(did string, routers []string) bool {
	idx, err := RouterLastIndexOf(did, routers)
	if err != nil {
		return false
	}
	return idx == len(routers)-1
}
