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
