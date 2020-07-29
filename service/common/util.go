package common

import (
	"encoding/json"
	"fmt"
	"git.ont.io/ontid/otf/common/packager"
	"git.ont.io/ontid/otf/utils"
	"io/ioutil"
	"strings"

	"git.ont.io/ontid/otf/common/log"
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

func ParseConnectionMsg(c *gin.Context, packager *ecdsa.Packager) (*message.Connection, *packager.MessageData, error) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		return nil, nil, err
	}
	msg, err := packager.UnPackData(body)
	if err != nil {
		return nil, nil, err
	}
	if msg.Connection == nil {
		return nil, msg.Message, nil
	}
	data, err := packager.UnPackConnection(msg)
	if err != nil {
		return nil, nil, err
	}
	connection := &message.Connection{}
	err = json.Unmarshal(data.Data, connection)
	if err != nil {
		return nil, nil, err
	}
	return connection, msg.Message, nil
}

func ParseMessage(enablePackage bool, ctx *gin.Context, packager *ecdsa.Packager, messageType MessageType, msgSvr *MsgService) (interface{}, bool, error) {
	msgObject, err := getMsgObjectByType(messageType)
	if err != nil {
		return nil, false, err
	}
	if enablePackage {
		connections, messageData, err := ParseConnectionMsg(ctx, packager)
		if err != nil {
			return nil, false, err
		}
		//check need router forward
		if connections != nil && !IsReceiver(msgSvr.Cfg.SelfDID, MergeRouter(connections.MyRouter, connections.TheirRouter)) {
			outMsg := OutboundMsg{
				Msg: Message{
					MessageType: TransferForwardMsgType(messageType),
					Content:     messageData,
				},
				Conn:      *connections,
				IsForward: true,
			}
			err = msgSvr.HandleOutBound(outMsg)
			if err != nil {
				log.Errorf("error on HandleOutBound:%s", err.Error())
				return nil, false, fmt.Errorf("handle forward msg error:%s", err)
			}
			return nil, true, nil
		}

		data, err := packager.UnpackMessage(messageData, msgSvr.Cfg.SelfDID)
		if err != nil {
			return nil, false, err
		}
		err = json.Unmarshal(data.Data, msgObject)
		if err != nil {
			return nil, false, err
		}
	} else {
		err = ctx.Bind(msgObject)
		if err != nil {
			return nil, false, err
		}

		connections := msgObject.GetConnection()
		if connections != nil {
			//check need router forward
			if !IsReceiver(msgSvr.Cfg.SelfDID, MergeRouter(connections.MyRouter, connections.TheirRouter)) {
				outMsg := OutboundMsg{
					Msg: Message{
						MessageType: TransferForwardMsgType(messageType),
						Content:     msgObject,
					},
					Conn:      *connections,
					IsForward: true,
				}
				err = msgSvr.HandleOutBound(outMsg)
				if err != nil {
					log.Errorf("error on HandleOutBound:%s", err.Error())
					return nil, false, fmt.Errorf("handle forward msg error:%s", err)
				}
				return nil, true, nil
			}
		}

	}
	return msgObject, false, nil
}

func getMsgObjectByType(messageType MessageType) (message.RequestInf, error) {
	var req message.RequestInf
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
	case SendBasicMsgType, ReceiveBasicMsgType:
		req = &message.BasicMessage{}
	case QueryBasicMessageType:
		req = &message.QueryBasicMessageRequest{}
	case QueryCredentialType:
		req = &message.QueryCredentialRequest{}
	case QueryPresentationType:
		req = &message.QueryPresentationRequest{}
	case DeleteCredentialType:
		req = &message.DeleteCredentialRequest{}
	case DeletePresentationType:
		req = &message.DeletePresentationRequest{}
	case QueryConnectionsType:
		req = &message.QueryConnectionsRequest{}
	default:
		return nil, fmt.Errorf("msg type err:%v", messageType)
	}
	return req, nil
}

func MergeRouter(myRouters, theirRouters []string) []string {
	return append(myRouters, reverseRouter(theirRouters)...)
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
