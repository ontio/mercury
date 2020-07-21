package controller

import (
	"fmt"
	"net/http"
	"strings"

	"git.ont.io/ontid/otf/common/log"
	"git.ont.io/ontid/otf/common/message"
	"git.ont.io/ontid/otf/common/packager/ecdsa"
	"git.ont.io/ontid/otf/service/common"
	"git.ont.io/ontid/otf/store"
	"git.ont.io/ontid/otf/utils"
	"git.ont.io/ontid/otf/vdri"
	"github.com/gin-gonic/gin"
)

type SystemController struct {
	packager *ecdsa.Packager
	store    store.Store
	msgSvr   *common.MsgService
}

func NewSystemController(packager *ecdsa.Packager, store store.Store,
	msgSvr *common.MsgService) common.Router {
	return &SystemController{
		packager: packager,
		store:    store,
		msgSvr:   msgSvr,
	}
}

func (c *SystemController) Routes() common.Routes {
	return common.Routes{
		{
			Name:        "Invitation",
			Method:      strings.ToUpper("Post"),
			Pattern:     "/api/v1/invitation",
			HandlerFunc: c.Invitation,
		},
		{
			Name:        "ConnectionRequest",
			Method:      strings.ToUpper("Post"),
			Pattern:     "/api/v1/connectionrequest",
			HandlerFunc: c.ConnectionRequest,
		},
		{
			Name:        "ConnectionResponse",
			Method:      strings.ToUpper("Post"),
			Pattern:     "/api/v1/connectionresponse",
			HandlerFunc: c.ConnectionResponse,
		},
		{
			Name:        "ConnectionAck",
			Method:      strings.ToUpper("Post"),
			Pattern:     "/api/v1/connectionack",
			HandlerFunc: c.ConnectionAck,
		},
		{
			Name:        "SendDisConnect",
			Method:      strings.ToUpper("Post"),
			Pattern:     "/api/v1/senddisconnect",
			HandlerFunc: c.SendDisConnect,
		},
		{
			Name:        "Disconnect",
			Method:      strings.ToUpper("Post"),
			Pattern:     "/api/v1/disconnect",
			HandlerFunc: c.Disconnect,
		},
		{
			Name:        "SendBasicMsg",
			Method:      strings.ToUpper("Post"),
			Pattern:     "/api/v1/sendbasicmsg",
			HandlerFunc: c.SendBasicMsg,
		},
		{
			Name:        "ReceiveBasicMsg",
			Method:      strings.ToUpper("Post"),
			Pattern:     "/api/v1/receivebasicmsg",
			HandlerFunc: c.ReceiveBasicMsg,
		},
		{
			Name:        "QueryBasicMsg",
			Method:      strings.ToUpper("Post"),
			Pattern:     "/api/v1/queryBasicMsg",
			HandlerFunc: c.QueryBasicMsg,
		},
	}
}

func (c *SystemController) Invitation(ctx *gin.Context) {
	resp := common.Gin{C: ctx}
	data, err := common.ParseMessage(common.EnablePackage, ctx, c.packager, message.InvitationType)
	if err != nil {
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
		return
	}
	invitation, ok := data.(*message.Invitation)
	if !ok {
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER,fmt.Errorf("data convert err").Error(), nil)
		return
	}
	err = c.SaveInvitation(*invitation)
	if err != nil {
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
		return
	}
	resp.Response(http.StatusOK, message.SUCCEED_CODE, "", invitation)
	return
}

func (c *SystemController) ConnectionRequest(ctx *gin.Context) {
	resp := common.Gin{C: ctx}
	data, err := common.ParseMessage(common.EnablePackage, ctx, c.packager, message.ConnectionRequestType)
	if err != nil {
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
		return
	}
	req, ok := data.(*message.ConnectionRequest)
	if !ok {
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER,fmt.Errorf("data convert err").Error(), nil)
		return
	}
	ivrc, err := c.GetInvitation(req.Connection.TheirDid, req.InvitationId)
	if err != nil {
		log.Infof("err on GetInvitation:%s\n", err.Error())
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
		return
	}
	err = c.SaveConnectionRequest(*req, message.ConnectionRequestReceived)
	if err != nil {
		log.Infof("err on SaveConnectionRequest:%s\n", err.Error())
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
		return
	}
	//update invitation to used state
	err = c.UpdateInvitation(ivrc.Invitation.Did, ivrc.Invitation.Id, message.InvitationUsed)
	if err != nil {
		log.Infof("err on UpdateInvitation:%s\n", err.Error())
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
		return
	}
	//send response outbound
	res := new(message.ConnectionResponse)
	res.Id = utils.GenUUID()
	res.Thread = message.Thread{
		ID: req.Id,
	}
	//todo define the response type
	res.Type = vdri.ConnectionResponseSpec
	//self conn
	res.Connection = message.Connection{
		MyDid:       ivrc.Invitation.Did,
		MyRouter:    ivrc.Invitation.Router,
		TheirDid:    req.Connection.MyDid,
		TheirRouter: req.Connection.MyRouter,
	}
	outMsg := message.Message{
		MessageType: message.ConnectionResponseType,
		Content:     res,
	}
	err = c.msgSvr.HandleOutBound(common.OutboundMsg{
		Msg:  outMsg,
		Conn: res.Connection,
	})
	if err != nil {
		log.Errorf("err on HandleOutBound:%s\n", err.Error())
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
		return
	}
	resp.Response(http.StatusOK, message.SUCCEED_CODE, "", nil)
	return
}

func (c *SystemController) ConnectionResponse(ctx *gin.Context) {
	resp := common.Gin{C: ctx}
	data, err := common.ParseMessage(common.EnablePackage, ctx, c.packager, message.ConnectionResponseType)
	if err != nil {
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
		return
	}
	req, ok := data.(*message.ConnectionResponse)
	if !ok {
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER,fmt.Errorf("data convert err").Error(), nil)
		return
	}
	connId := req.Thread.ID
	//2. create and save a connection object
	err = c.SaveConnection(common.ReverseConnection(req.Connection))
	if err != nil {
		log.Errorf("err on SaveConnection:%s\n", err.Error())
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
		return
	}
	//3. send ACK back
	ack := message.ConnectionACK{
		Type:       vdri.ConnectionACKSpec,
		Id:         utils.GenUUID(),
		Thread:     message.Thread{ID: connId},
		Status:     utils.ACK_SUCCEED,
		Connection: common.ReverseConnection(req.Connection),
	}
	outMsg := message.Message{
		MessageType: message.ConnectionACKType,
		Content:     ack,
	}
	err = c.msgSvr.HandleOutBound(common.OutboundMsg{
		Msg:  outMsg,
		Conn: ack.Connection,
	})
	if err != nil {
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
		return
	}
	resp.Response(http.StatusOK, message.SUCCEED_CODE, "", nil)
	return
}

func (c *SystemController) ConnectionAck(ctx *gin.Context) {
	resp := common.Gin{C: ctx}
	data, err := common.ParseMessage(common.EnablePackage, ctx, c.packager, message.ConnectionACKType)
	if err != nil {
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
		return
	}
	req, ok := data.(*message.ConnectionACK)
	if !ok {
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER,fmt.Errorf("data convert err").Error(), nil)
		return
	}
	if req.Status != utils.ACK_SUCCEED {
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, fmt.Errorf("got failed ACK ").Error(), nil)
		return
	}
	connId := req.Thread.ID
	err = c.UpdateConnectionRequest(req.Connection.TheirDid, connId, message.ConnectionACKReceived)
	if err != nil {
		log.Errorf("err on UpdateConnectionRequest:%s\n", err.Error())
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
		return
	}
	cr, err := c.GetConnectionRequest(req.Connection.TheirDid, connId)
	if err != nil {
		log.Errorf("err on GetConnectionRequest:%s\n", err.Error())
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
		return
	}
	err = c.SaveConnection(common.ReverseConnection(cr.ConnReq.Connection))
	if err != nil {
		log.Errorf("err on SaveConnection:%s\n", err.Error())
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
		return
	}
	resp.Response(http.StatusOK, message.SUCCEED_CODE, "", nil)
	return
}

func (c *SystemController) SendDisConnect(ctx *gin.Context) {
	resp := common.Gin{C: ctx}
	data, err := common.ParseMessage(common.EnablePackage, ctx, c.packager, message.SendDisconnectType)
	if err != nil {
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
		return
	}
	req, ok := data.(*message.DisconnectRequest)
	if !ok {
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER,fmt.Errorf("data convert err").Error(), nil)
		return
	}
	myDid := req.Connection.MyDid
	theirDid := req.Connection.TheirDid
	//1. remove connection
	err = c.DeleteConnection(myDid, theirDid)
	if err != nil {
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
		return
	}
	outMsg := message.Message{
		MessageType: message.DisconnectType,
		Content:     req,
	}
	err = c.msgSvr.HandleOutBound(common.OutboundMsg{
		Msg:  outMsg,
		Conn: common.ReverseConnection(req.Connection),
	})
	if err != nil {
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
		return
	}
	resp.Response(http.StatusOK, message.SUCCEED_CODE, "", nil)
	return
}

func (c *SystemController) Disconnect(ctx *gin.Context) {
	resp := common.Gin{C: ctx}
	data, err := common.ParseMessage(common.EnablePackage, ctx, c.packager, message.DisconnectType)
	if err != nil {
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
		return
	}
	req, ok := data.(*message.DisconnectRequest)
	if !ok {
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER,fmt.Errorf("data convert err").Error(), nil)
		return
	}
	err = c.DeleteConnection(req.Connection.TheirDid, req.Connection.MyDid)
	if err != nil {
		log.Errorf("error:%s", err.Error())
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
		return
	}
	resp.Response(http.StatusOK, message.SUCCEED_CODE, "", nil)
	return
}

func (c *SystemController) SendBasicMsg(ctx *gin.Context) {
	resp := common.Gin{C: ctx}
	data, err := common.ParseMessage(common.EnablePackage, ctx, c.packager, message.SendBasicMsgType)
	if err != nil {
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
		return
	}
	req, ok := data.(*message.BasicMessage)
	if !ok {
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER,fmt.Errorf("data convert err").Error(), nil)
		return
	}
	conn, err := c.GetConnection(req.Connection.MyDid, req.Connection.TheirDid)
	if err != nil {
		log.Errorf("err on GetConnection:%s\n", err.Error())
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
		return
	}
	req.Type = vdri.BasicMsgSpec
	req.Id = utils.GenUUID()
	outMsg := common.OutboundMsg{
		Msg: message.Message{
			MessageType: message.ReceiveBasicMsgType,
			Content:     req,
		},
		Conn: conn,
	}
	err = c.msgSvr.HandleOutBound(outMsg)
	if err != nil {
		log.Errorf("err on HandleOutBound:%s\n", err.Error())
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
		return
	}
	err = c.SaveGeneralMsg(req, true)
	if err != nil {
		log.Errorf("err on HandleOutBound:%s\n", err.Error())
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
		return
	}
	resp.Response(http.StatusOK, message.SUCCEED_CODE, "", nil)
	return
}

func (c *SystemController) ReceiveBasicMsg(ctx *gin.Context) {
	resp := common.Gin{C: ctx}
	data, err := common.ParseMessage(common.EnablePackage, ctx, c.packager, message.SendBasicMsgType)
	if err != nil {
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
		return
	}
	req, ok := data.(*message.BasicMessage)
	if !ok {
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER,fmt.Errorf("data convert err").Error(), nil)
		return
	}
	err = utils.CheckConnection(req.Connection.TheirDid, req.Connection.MyDid, c.store)
	if err != nil {
		log.Infof("no connect found with did:%s", req.Connection.MyDid)
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
		return
	}
	err = c.SaveGeneralMsg(req, false)
	if err != nil {
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
		return
	}
	resp.Response(http.StatusOK, message.SUCCEED_CODE, "", nil)
	return
}

func (c *SystemController) QueryBasicMsg(ctx *gin.Context) {
	resp := common.Gin{C: ctx}
	data, err := common.ParseMessage(common.EnablePackage, ctx, c.packager, message.QueryBasicMessageType)
	if err != nil {
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
		return
	}
	req, ok := data.(*message.QueryGeneralMessageRequest)
	if !ok {
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER,fmt.Errorf("data convert err").Error(), nil)
		return
	}
	ret, err := c.QueryGeneraMsg(req.DID, req.Latest, req.RemoveAfterRead)
	if err != nil {
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
		return
	}
	resp.Response(http.StatusOK, message.SUCCEED_CODE, "", ret)
	return
}
