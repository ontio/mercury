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

func (s *SystemController) Routes() common.Routes {
	return common.Routes{
		{
			Name:        "Invitation",
			Method:      strings.ToUpper("Post"),
			Pattern:     common.InviteApi,
			HandlerFunc: s.Invitation,
		},
		{
			Name:        "ConnectionRequest",
			Method:      strings.ToUpper("Post"),
			Pattern:     common.ConnectRequestApi,
			HandlerFunc: s.ConnectionRequest,
		},
		{
			Name:        "ConnectionResponse",
			Method:      strings.ToUpper("Post"),
			Pattern:     common.ConnectResponseApi,
			HandlerFunc: s.ConnectionResponse,
		},
		{
			Name:        "ConnectionAck",
			Method:      strings.ToUpper("Post"),
			Pattern:     common.ConnectAckApi,
			HandlerFunc: s.ConnectionAck,
		},
		{
			Name:        "SendDisConnect",
			Method:      strings.ToUpper("Post"),
			Pattern:     common.SendDisconnectApi,
			HandlerFunc: s.SendDisConnect,
		},
		{
			Name:        "Disconnect",
			Method:      strings.ToUpper("Post"),
			Pattern:     common.DisconnectApi,
			HandlerFunc: s.Disconnect,
		},
		{
			Name:        "SendBasicMsg",
			Method:      strings.ToUpper("Post"),
			Pattern:     common.SendBasicMsgApi,
			HandlerFunc: s.SendBasicMsg,
		},
		{
			Name:        "ReceiveBasicMsg",
			Method:      strings.ToUpper("Post"),
			Pattern:     common.ReceiveBasicMsgApi,
			HandlerFunc: s.ReceiveBasicMsg,
		},
		{
			Name:        "QueryBasicMsg",
			Method:      strings.ToUpper("Post"),
			Pattern:     common.QueryBasicMsgApi,
			HandlerFunc: s.QueryBasicMsg,
		},
	}
}

func (s *SystemController) Invitation(ctx *gin.Context) {
	resp := common.Gin{C: ctx}
	data, isForward, err := common.ParseMessage(common.EnablePackage, ctx, s.packager, common.InvitationType, s.msgSvr)
	if err != nil {
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
		return
	}
	if isForward {
		resp.Response(http.StatusOK, message.SUCCEED_CODE, "", nil)
		return
	}
	invitation, ok := data.(*message.Invitation)
	if !ok {
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, fmt.Errorf("data convert err").Error(), nil)
		return
	}
	err = s.SaveInvitation(*invitation)
	if err != nil {
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
		return
	}
	resp.Response(http.StatusOK, message.SUCCEED_CODE, "", invitation)
	return
}

func (s *SystemController) ConnectionRequest(ctx *gin.Context) {
	resp := common.Gin{C: ctx}
	data, isForward, err := common.ParseMessage(common.EnablePackage, ctx, s.packager, common.ConnectionRequestType, s.msgSvr)
	if err != nil {
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
		return
	}
	if isForward {
		resp.Response(http.StatusOK, message.SUCCEED_CODE, "", nil)
		return
	}
	req, ok := data.(*message.ConnectionRequest)
	if !ok {
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, fmt.Errorf("data convert err").Error(), nil)
		return
	}
	ivrc, err := s.GetInvitation(req.Connection.TheirDid, req.InvitationId)
	if err != nil {
		log.Infof("err on GetInvitation:%s\n", err.Error())
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
		return
	}
	err = s.SaveConnectionRequest(*req, message.ConnectionRequestReceived)
	if err != nil {
		log.Infof("err on SaveConnectionRequest:%s\n", err.Error())
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
		return
	}
	//update invitation to used state
	err = s.UpdateInvitation(ivrc.Invitation.Did, ivrc.Invitation.Id, message.InvitationUsed)
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
	outMsg := common.Message{
		MessageType: common.ConnectionResponseType,
		Content:     res,
	}
	err = s.msgSvr.HandleOutBound(common.OutboundMsg{
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

func (s *SystemController) ConnectionResponse(ctx *gin.Context) {
	resp := common.Gin{C: ctx}
	data, isForward, err := common.ParseMessage(common.EnablePackage, ctx, s.packager, common.ConnectionResponseType, s.msgSvr)
	if err != nil {
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
		return
	}
	if isForward {
		resp.Response(http.StatusOK, message.SUCCEED_CODE, "", nil)
		return
	}
	req, ok := data.(*message.ConnectionResponse)
	if !ok {
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, fmt.Errorf("data convert err").Error(), nil)
		return
	}
	connId := req.Thread.ID
	err = s.SaveConnection(common.ReverseConnection(req.Connection))
	if err != nil {
		log.Errorf("err on SaveConnection:%s\n", err.Error())
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
		return
	}
	ack := message.ConnectionACK{
		Type:       vdri.ConnectionACKSpec,
		Id:         utils.GenUUID(),
		Thread:     message.Thread{ID: connId},
		Status:     utils.ACK_SUCCEED,
		Connection: common.ReverseConnection(req.Connection),
	}
	outMsg := common.Message{
		MessageType: common.ConnectionAckType,
		Content:     ack,
	}
	err = s.msgSvr.HandleOutBound(common.OutboundMsg{
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

func (s *SystemController) ConnectionAck(ctx *gin.Context) {
	resp := common.Gin{C: ctx}
	data, isForward, err := common.ParseMessage(common.EnablePackage, ctx, s.packager, common.ConnectionAckType, s.msgSvr)
	if err != nil {
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
		return
	}
	if isForward {
		resp.Response(http.StatusOK, message.SUCCEED_CODE, "", nil)
		return
	}
	req, ok := data.(*message.ConnectionACK)
	if !ok {
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, fmt.Errorf("data convert err").Error(), nil)
		return
	}
	if req.Status != utils.ACK_SUCCEED {
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, fmt.Errorf("got failed ACK ").Error(), nil)
		return
	}
	connId := req.Thread.ID
	err = s.UpdateConnectionRequest(req.Connection.TheirDid, connId, message.ConnectionACKReceived)
	if err != nil {
		log.Errorf("err on UpdateConnectionRequest:%s\n", err.Error())
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
		return
	}
	cr, err := s.GetConnectionRequest(req.Connection.TheirDid, connId)
	if err != nil {
		log.Errorf("err on GetConnectionRequest:%s\n", err.Error())
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
		return
	}
	err = s.SaveConnection(common.ReverseConnection(cr.ConnReq.Connection))
	if err != nil {
		log.Errorf("err on SaveConnection:%s\n", err.Error())
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
		return
	}
	resp.Response(http.StatusOK, message.SUCCEED_CODE, "", nil)
	return
}

func (s *SystemController) SendDisConnect(ctx *gin.Context) {
	resp := common.Gin{C: ctx}
	data, isForward, err := common.ParseMessage(common.EnablePackage, ctx, s.packager, common.SendDisconnectType, s.msgSvr)
	if err != nil {
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
		return
	}
	if isForward {
		resp.Response(http.StatusOK, message.SUCCEED_CODE, "", nil)
		return
	}
	req, ok := data.(*message.DisconnectRequest)
	if !ok {
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, fmt.Errorf("data convert err").Error(), nil)
		return
	}

	//check connection
	_, err = s.GetConnection(req.Connection.MyDid, req.Connection.TheirDid)
	if err != nil {
		log.Errorf("err on GetConnection:%s\n", err.Error())
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
		return
	}

	myDid := req.Connection.MyDid
	theirDid := req.Connection.TheirDid
	err = s.DeleteConnection(myDid, theirDid)
	if err != nil {
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
		return
	}
	outMsg := common.Message{
		MessageType: common.DisconnectType,
		Content:     req,
	}
	err = s.msgSvr.HandleOutBound(common.OutboundMsg{
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

func (s *SystemController) Disconnect(ctx *gin.Context) {
	resp := common.Gin{C: ctx}
	data, isForward, err := common.ParseMessage(common.EnablePackage, ctx, s.packager, common.DisconnectType, s.msgSvr)
	if err != nil {
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
		return
	}
	if isForward {
		resp.Response(http.StatusOK, message.SUCCEED_CODE, "", nil)
		return
	}
	req, ok := data.(*message.DisconnectRequest)
	if !ok {
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, fmt.Errorf("data convert err").Error(), nil)
		return
	}

	err = s.DeleteConnection(req.Connection.TheirDid, req.Connection.MyDid)
	if err != nil {
		log.Errorf("error:%s", err.Error())
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
		return
	}
	resp.Response(http.StatusOK, message.SUCCEED_CODE, "", nil)
	return
}

func (s *SystemController) SendBasicMsg(ctx *gin.Context) {
	resp := common.Gin{C: ctx}
	data, isForward, err := common.ParseMessage(common.EnablePackage, ctx, s.packager, common.SendBasicMsgType, s.msgSvr)
	if err != nil {
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
		return
	}
	if isForward {
		resp.Response(http.StatusOK, message.SUCCEED_CODE, "", nil)
		return
	}
	req, ok := data.(*message.BasicMessage)
	if !ok {
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, fmt.Errorf("data convert err").Error(), nil)
		return
	}
	conn, err := s.GetConnection(req.Connection.MyDid, req.Connection.TheirDid)
	if err != nil {
		log.Errorf("err on GetConnection:%s\n", err.Error())
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
		return
	}
	req.Type = vdri.BasicMsgSpec
	req.Id = utils.GenUUID()
	outMsg := common.OutboundMsg{
		Msg: common.Message{
			MessageType: common.ReceiveBasicMsgType,
			Content:     req,
		},
		Conn: conn,
	}
	err = s.msgSvr.HandleOutBound(outMsg)
	if err != nil {
		log.Errorf("err on HandleOutBound:%s\n", err.Error())
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
		return
	}
	err = s.SaveGeneralMsg(req, true)
	if err != nil {
		log.Errorf("err on HandleOutBound:%s\n", err.Error())
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
		return
	}
	resp.Response(http.StatusOK, message.SUCCEED_CODE, "", nil)
	return
}

func (s *SystemController) ReceiveBasicMsg(ctx *gin.Context) {
	resp := common.Gin{C: ctx}
	data, isForward, err := common.ParseMessage(common.EnablePackage, ctx, s.packager, common.SendBasicMsgType, s.msgSvr)
	if err != nil {
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
		return
	}
	if isForward {
		resp.Response(http.StatusOK, message.SUCCEED_CODE, "", nil)
		return
	}
	req, ok := data.(*message.BasicMessage)
	if !ok {
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, fmt.Errorf("data convert err").Error(), nil)
		return
	}
	err = utils.CheckConnection(req.Connection.TheirDid, req.Connection.MyDid, s.store)
	if err != nil {
		log.Infof("no connect found with did:%s", req.Connection.MyDid)
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
		return
	}
	err = s.SaveGeneralMsg(req, false)
	if err != nil {
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
		return
	}
	resp.Response(http.StatusOK, message.SUCCEED_CODE, "", nil)
	return
}

func (s *SystemController) QueryBasicMsg(ctx *gin.Context) {
	resp := common.Gin{C: ctx}
	data, isForward, err := common.ParseMessage(common.EnablePackage, ctx, s.packager, common.QueryBasicMessageType, s.msgSvr)
	if err != nil {
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
		return
	}
	if isForward {
		resp.Response(http.StatusOK, message.SUCCEED_CODE, "", nil)
		return
	}
	req, ok := data.(*message.QueryGeneralMessageRequest)
	if !ok {
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, fmt.Errorf("data convert err").Error(), nil)
		return
	}
	ret, err := s.QueryGeneraMsg(req.DID, req.Latest, req.RemoveAfterRead)
	if err != nil {
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
		return
	}
	resp.Response(http.StatusOK, message.SUCCEED_CODE, "", ret)
	return
}
