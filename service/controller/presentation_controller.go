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

type PresentationController struct {
	packager *ecdsa.Packager
	store    store.Store
	msgSvr   *common.MsgService
	vdri     vdri.VDRI
}

func NewPresentationController(packager *ecdsa.Packager, store store.Store,
	msgSvr *common.MsgService, v vdri.VDRI) common.Router {
	return &PresentationController{
		packager: packager,
		store:    store,
		msgSvr:   msgSvr,
		vdri:     v,
	}
}

func (c *PresentationController) Routes() common.Routes {
	return common.Routes{
		{
			Name:        "SendRequestPresentation",
			Method:      strings.ToUpper("Post"),
			Pattern:     common.SendRequestPresentationApi,
			HandlerFunc: c.SendRequestPresentation,
		},
		{
			Name:        "RequestPresentation",
			Method:      strings.ToUpper("Post"),
			Pattern:     common.RequestPresentationApi,
			HandlerFunc: c.RequestPresentation,
		},
		{
			Name:        "PresentProof",
			Method:      strings.ToUpper("Post"),
			Pattern:     common.PresentationProofApi,
			HandlerFunc: c.PresentProof,
		},
		{
			Name:        "PresentationAck",
			Method:      strings.ToUpper("Post"),
			Pattern:     common.PresentationAckApi,
			HandlerFunc: c.PresentationAck,
		},
		{
			Name:        "QueryPresentation",
			Method:      strings.ToUpper("Post"),
			Pattern:     common.QueryPresentationApi,
			HandlerFunc: c.QueryPresentation,
		},
		{
			Name:        "DeletePresentation",
			Method:      strings.ToUpper("Post"),
			Pattern:     common.DeletePresentationApi,
			HandlerFunc: c.DeletePresentation,
		},
	}
}

func (p *PresentationController) SendRequestPresentation(ctx *gin.Context) {
	resp := common.Gin{C: ctx}
	data, isForward, err := common.ParseMessage(common.EnablePackage, ctx, p.packager, common.SendRequestPresentationType, p.msgSvr)
	if err != nil {
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
		return
	}
	if isForward {
		resp.Response(http.StatusOK, message.SUCCEED_CODE, "", nil)
		return
	}
	req, ok := data.(*message.RequestPresentation)
	if !ok {
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, fmt.Errorf("data convert err").Error(), nil)
		return
	}

	err = utils.CheckConnection(req.Connection.MyDid, req.Connection.TheirDid, p.store)
	if err != nil {
		log.Errorf("no connect found with did:%s", req.Connection.MyDid)
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
		return
	}

	outMsg := common.OutboundMsg{
		Msg: common.Message{
			MessageType: common.RequestPresentationType,
			Content:     req,
		},
		Conn: req.Connection,
	}
	err = p.msgSvr.HandleOutBound(outMsg)
	if err != nil {
		log.Errorf("error on HandleOutBound :%s", err.Error())
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
		return
	}
	resp.Response(http.StatusOK, message.SUCCEED_CODE, "", nil)
	return
}

func (p *PresentationController) RequestPresentation(ctx *gin.Context) {
	resp := common.Gin{C: ctx}
	data, isForward, err := common.ParseMessage(common.EnablePackage, ctx, p.packager, common.RequestPresentationType, p.msgSvr)
	if err != nil {
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
		return
	}
	if isForward {
		resp.Response(http.StatusOK, message.SUCCEED_CODE, "", nil)
		return
	}
	req, ok := data.(*message.RequestPresentation)
	if !ok {
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, fmt.Errorf("data convert err").Error(), nil)
		return
	}

	err = utils.CheckConnection(req.Connection.TheirDid, req.Connection.MyDid, p.store)
	if err != nil {
		log.Infof("no connect found with did:%s", req.Connection.MyDid)
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
		return
	}
	presentation, err := p.vdri.PresentProof(req, p.store)
	if err != nil {
		log.Errorf("errors on PresentProof :%s", err.Error())
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
		return
	}

	err = p.SaveRequestPresentation(req.Connection.MyDid, req.Id, *req)
	if err != nil {
		log.Errorf("error on SaveRequestPresentation:%s", err.Error())
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
		return
	}
	outMsg := common.OutboundMsg{
		Msg: common.Message{
			MessageType: common.PresentationType,
			Content:     presentation,
		},
		Conn: common.ReverseConnection(presentation.Connection),
	}
	err = p.msgSvr.HandleOutBound(outMsg)
	if err != nil {
		log.Errorf("error on HandleOutBound:%s", err.Error())
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
		return
	}
	resp.Response(http.StatusOK, message.SUCCEED_CODE, "", nil)
	return
}

func (p *PresentationController) PresentProof(ctx *gin.Context) {
	resp := common.Gin{C: ctx}
	data, isForward, err := common.ParseMessage(common.EnablePackage, ctx, p.packager, common.PresentationType, p.msgSvr)
	if err != nil {
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
		return
	}
	if isForward {
		resp.Response(http.StatusOK, message.SUCCEED_CODE, "", nil)
		return
	}
	req, ok := data.(*message.Presentation)
	if !ok {
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, fmt.Errorf("data convert err").Error(), nil)
		return
	}

	err = utils.CheckConnection(req.Connection.TheirDid, req.Connection.MyDid, p.store)
	if err != nil {
		log.Infof("no connect found with did:%s", req.Connection.MyDid)
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
		return
	}
	err = p.SavePresentation(req.Connection.TheirDid, req.Thread.ID, *req)
	if err != nil {
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
		return
	}
	ack := &message.PresentationACK{
		Id:         utils.GenUUID(),
		Thread:     req.Thread,
		Connection: common.ReverseConnection(req.Connection),
		Type:       vdri.PresentationACKSpec,
		Status:     utils.ACK_SUCCEED,
	}
	outMsg := common.OutboundMsg{
		Msg: common.Message{
			MessageType: common.PresentationAckType,
			Content:     ack,
		},
		Conn: ack.Connection,
	}
	err = p.msgSvr.HandleOutBound(outMsg)
	if err != nil {
		log.Errorf("error on HandleOutBound:%s", err.Error())
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
		return
	}
	resp.Response(http.StatusOK, message.SUCCEED_CODE, "", nil)
	return
}

func (p *PresentationController) PresentationAck(ctx *gin.Context) {
	resp := common.Gin{C: ctx}
	data, isForward, err := common.ParseMessage(common.EnablePackage, ctx, p.packager, common.PresentationAckType, p.msgSvr)
	if err != nil {
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
		return
	}
	if isForward {
		resp.Response(http.StatusOK, message.SUCCEED_CODE, "", nil)
		return
	}
	req, ok := data.(*message.PresentationACK)
	if !ok {
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, fmt.Errorf("data convert err").Error(), nil)
		return
	}
	err = utils.CheckConnection(req.Connection.TheirDid, req.Connection.MyDid, p.store)
	if err != nil {
		log.Infof("no connect found with did:%s", req.Connection.MyDid)
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
		return
	}
	err = p.UpdateRequestPresentaion(req.Connection.MyDid, req.Thread.ID, message.RequestPresentationReceived)
	if err != nil {
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
		return
	}
	resp.Response(http.StatusOK, message.SUCCEED_CODE, "", nil)
	return
}

func (p *PresentationController) QueryPresentation(ctx *gin.Context) {
	resp := common.Gin{C: ctx}
	data, isForward, err := common.ParseMessage(common.EnablePackage, ctx, p.packager, common.QueryPresentationType, p.msgSvr)
	if err != nil {
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
		return
	}
	if isForward {
		resp.Response(http.StatusOK, message.SUCCEED_CODE, "", nil)
		return
	}
	req, ok := data.(*message.QueryPresentationRequest)
	if !ok {
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, fmt.Errorf("data convert err").Error(), nil)
		return
	}
	rec, err := p.QueryPresentationFromStore(req.DId, req.Id)
	if err != nil {
		log.Errorf("error on QueryPresentationType:%s", err.Error())
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
		return
	}
	resp.Response(http.StatusOK, message.SUCCEED_CODE, "", &message.QueryPresentationResponse{
		Formats:            rec.Formats,
		PresentationAttach: rec.PresentationAttach,
	})
	return
}
func (c *PresentationController) DeletePresentation(ctx *gin.Context) {
	resp := common.Gin{C: ctx}
	data, _, err := common.ParseMessage(common.EnablePackage, ctx, c.packager, common.DeletePresentationType, c.msgSvr)
	if err != nil {
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
		return
	}
	req, ok := data.(*message.DeletePresentationRequest)
	if !ok {
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, fmt.Errorf("data convert err").Error(), nil)
		return
	}
	err = c.DelPresentation(req.DId, req.Id)
	if err != nil {
		log.Errorf("error on QueryPresentationType:%s", err.Error())
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
		return
	}
	resp.Response(http.StatusOK, message.SUCCEED_CODE, "", nil)
	return
}
