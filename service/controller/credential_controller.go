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

type CredentialController struct {
	packager *ecdsa.Packager
	store    store.Store
	msgSvr   *common.MsgService
	vdri     vdri.VDRI
}

func NewCredentialController(packager *ecdsa.Packager, store store.Store,
	msgSvr *common.MsgService, v vdri.VDRI) common.Router {
	return &CredentialController{
		packager: packager,
		store:    store,
		msgSvr:   msgSvr,
		vdri:     v,
	}
}

func (c *CredentialController) Routes() common.Routes {
	return common.Routes{
		{
			Name:        "SendProposalCredential",
			Method:      strings.ToUpper("Post"),
			Pattern:     common.SendProposalCredentialReqApi,
			HandlerFunc: c.SendProposalCredential,
		},
		{
			Name:        "ProposalCredential",
			Method:      strings.ToUpper("Post"),
			Pattern:     common.ProposalCredentialReqApi,
			HandlerFunc: c.ProposalCredential,
		},
		{
			Name:        "OfferCredential",
			Method:      strings.ToUpper("Post"),
			Pattern:     common.OfferCredentialApi,
			HandlerFunc: c.OfferCredential,
		},
		{
			Name:        "SendRequestCredential",
			Method:      strings.ToUpper("Post"),
			Pattern:     common.SendRequestCredentialApi,
			HandlerFunc: c.SendRequestCredential,
		},
		{
			Name:        "RequestCredential",
			Method:      strings.ToUpper("Post"),
			Pattern:     common.RequestCredentialApi,
			HandlerFunc: c.RequestCredential,
		},
		{
			Name:        "IssueCredential",
			Method:      strings.ToUpper("Post"),
			Pattern:     common.IssueCredentialApi,
			HandlerFunc: c.IssueCredential,
		},
		{
			Name:        "CredentialAck",
			Method:      strings.ToUpper("Post"),
			Pattern:     common.CredentialAckApi,
			HandlerFunc: c.CredentialAck,
		},
		{
			Name:        "QueryCredentail",
			Method:      strings.ToUpper("Post"),
			Pattern:     common.QueryCredentialApi,
			HandlerFunc: c.QueryCredential,
		},
	}
}

func (c *CredentialController) SendProposalCredential(ctx *gin.Context) {
	resp := common.Gin{C: ctx}
	data, isForward, err := common.ParseMessage(common.EnablePackage, ctx, c.packager, common.SendProposalCredentialType, c.msgSvr)
	if err != nil {
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
		return
	}
	if isForward {
		resp.Response(http.StatusOK, message.SUCCEED_CODE, "", nil)
		return
	}
	req, ok := data.(*message.ProposalCredential)
	if !ok {
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, fmt.Errorf("data convert err").Error(), nil)
		return
	}

	err = utils.CheckConnection(req.Connection.MyDid, req.Connection.TheirDid, c.store)
	if err != nil {
		log.Infof("no connect found with did:%s", req.Connection.MyDid)
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
		return
	}

	outMsg := common.OutboundMsg{
		Msg: common.Message{
			MessageType: common.ProposalCredentialType,
			Content:     req,
		},
		Conn: req.Connection,
	}
	err = c.msgSvr.HandleOutBound(outMsg)
	if err != nil {
		log.Errorf("error on HandleOutBound:%s", err.Error())
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
		return
	}
	resp.Response(http.StatusOK, message.SUCCEED_CODE, "", nil)
	return
}

func (c *CredentialController) ProposalCredential(ctx *gin.Context) {
	resp := common.Gin{C: ctx}
	data, isForward, err := common.ParseMessage(common.EnablePackage, ctx, c.packager, common.ProposalCredentialType, c.msgSvr)
	if err != nil {
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
		return
	}
	if isForward {
		resp.Response(http.StatusOK, message.SUCCEED_CODE, "", nil)
		return
	}
	req, ok := data.(*message.ProposalCredential)
	if !ok {
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, fmt.Errorf("data convert err").Error(), nil)
		return
	}
	err = utils.CheckConnection(req.Connection.TheirDid, req.Connection.MyDid, c.store)
	if err != nil {
		log.Infof("no connect found with did:%s", req.Connection.MyDid)
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
		return
	}
	log.Infof("proposal is %v", req)
	offer, err := c.vdri.OfferCredential(req)
	if err != nil {
		log.Errorf("error on offerCredetial")
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
		return
	}
	outerMsg := common.OutboundMsg{
		Msg: common.Message{
			MessageType: common.OfferCredentialType,
			Content:     offer,
		},
		Conn: offer.Connection,
	}
	err = c.msgSvr.HandleOutBound(outerMsg)
	if err != nil {
		log.Errorf("error on HandleOutBound :%s", err.Error())
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
		return
	}
	resp.Response(http.StatusOK, message.SUCCEED_CODE, "", nil)
	return
}

func (c *CredentialController) OfferCredential(ctx *gin.Context) {
	resp := common.Gin{C: ctx}
	data, isForward, err := common.ParseMessage(common.EnablePackage, ctx, c.packager, common.OfferCredentialType, c.msgSvr)
	if err != nil {
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
		return
	}
	if isForward {
		resp.Response(http.StatusOK, message.SUCCEED_CODE, "", nil)
		return
	}
	req, ok := data.(*message.OfferCredential)
	if !ok {
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, fmt.Errorf("data convert err").Error(), nil)
		return
	}
	err = utils.CheckConnection(req.Connection.TheirDid, req.Connection.MyDid, c.store)
	if err != nil {
		log.Infof("no connect found with did:%s", req.Connection.MyDid)
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
		return
	}
	err = c.SaveOfferCredential(req.Connection.TheirDid, req.Thread.ID, req)
	if err != nil {
		log.Errorf("error on SaveOfferCredential:%s", err.Error())
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
		return
	}
	resp.Response(http.StatusOK, message.SUCCEED_CODE, "", nil)
	return
}

func (c *CredentialController) SendRequestCredential(ctx *gin.Context) {
	resp := common.Gin{C: ctx}
	data, isForward, err := common.ParseMessage(common.EnablePackage, ctx, c.packager, common.RequestCredentialType, c.msgSvr)
	if err != nil {
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
		return
	}
	if isForward {
		resp.Response(http.StatusOK, message.SUCCEED_CODE, "", nil)
		return
	}
	req, ok := data.(*message.RequestCredential)
	if !ok {
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, fmt.Errorf("data convert err").Error(), nil)
		return
	}

	log.Infof("========before checkConnection ==========")
	err = utils.CheckConnection(req.Connection.MyDid, req.Connection.TheirDid, c.store)
	if err != nil {
		log.Errorf("no connect found with did:%s", req.Connection.MyDid)
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
		return
	}
	log.Infof("========after checkConnection ==========")

	outMsg := common.OutboundMsg{
		Msg: common.Message{
			MessageType: common.RequestCredentialType,
			Content:     req,
		},
		Conn: req.Connection,
	}
	err = c.msgSvr.HandleOutBound(outMsg)
	if err != nil {
		log.Errorf("error on HandleOutBound:%s", err.Error())
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
		return
	}
	resp.Response(http.StatusOK, message.SUCCEED_CODE, "", nil)
	return
}

func (c *CredentialController) RequestCredential(ctx *gin.Context) {
	resp := common.Gin{C: ctx}
	data, isForward, err := common.ParseMessage(common.EnablePackage, ctx, c.packager, common.RequestCredentialType, c.msgSvr)
	if err != nil {
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
		return
	}
	if isForward {
		resp.Response(http.StatusOK, message.SUCCEED_CODE, "", nil)
		return
	}
	req, ok := data.(*message.RequestCredential)
	if !ok {
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, fmt.Errorf("data convert err").Error(), nil)
		return
	}
	err = utils.CheckConnection(req.Connection.TheirDid, req.Connection.MyDid, c.store)
	if err != nil {
		log.Infof("no connect found with did:%s", req.Connection.MyDid)
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
		return
	}
	err = c.SaveRequestCredential(req.Connection.MyDid, req.Id, *req)
	if err != nil {
		log.Errorf("error on SaveRequestCredential:%s\n", err.Error())
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
		return
	}

	credential, err := c.vdri.IssueCredential(req)
	if err != nil {
		log.Errorf("error on IssueCredential:%s\n", err.Error())
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
		return
	}
	outMsg := common.OutboundMsg{
		Msg: common.Message{
			MessageType: common.IssueCredentialType,
			Content:     credential,
		},
		Conn: credential.Connection,
	}
	err = c.msgSvr.HandleOutBound(outMsg)
	if err != nil {
		log.Errorf("error on HandleOutBound:%s\n", err.Error())
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
		return
	}
	resp.Response(http.StatusOK, message.SUCCEED_CODE, "", nil)
	return
}

func (c *CredentialController) IssueCredential(ctx *gin.Context) {
	resp := common.Gin{C: ctx}
	data, isForward, err := common.ParseMessage(common.EnablePackage, ctx, c.packager, common.IssueCredentialType, c.msgSvr)
	if err != nil {
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
		return
	}
	if isForward {
		resp.Response(http.StatusOK, message.SUCCEED_CODE, "", nil)
		return
	}
	req, ok := data.(*message.IssueCredential)
	if !ok {
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, fmt.Errorf("data convert err").Error(), nil)
		return
	}

	err = c.SaveCredential(req.Connection.TheirDid, req.Thread.ID, *req)
	if err != nil {
		log.Errorf("error on SaveCredential:%s\n", err.Error())
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
		return
	}
	ack := message.CredentialACK{
		Type: vdri.CredentialACKSpec,
		Id:   utils.GenUUID(),
		Thread: message.Thread{
			ID: req.Thread.ID,
		},
		Status:     utils.ACK_SUCCEED,
		Connection: common.ReverseConnection(req.Connection),
	}
	outMsg := common.OutboundMsg{
		Msg: common.Message{
			MessageType: common.CredentialAckType,
			Content:     ack,
		},
		Conn: ack.Connection,
	}
	err = c.msgSvr.HandleOutBound(outMsg)
	if err != nil {
		log.Errorf("error on SaveCredential:%s\n", err.Error())
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
		return
	}
	resp.Response(http.StatusOK, message.SUCCEED_CODE, "", nil)
	return
}

func (c *CredentialController) CredentialAck(ctx *gin.Context) {
	resp := common.Gin{C: ctx}
	data, isForward, err := common.ParseMessage(common.EnablePackage, ctx, c.packager, common.ConnectionAckType, c.msgSvr)
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

	err = c.UpdateRequestCredential(req.Connection.MyDid, req.Thread.ID, message.RequestCredentialResolved)
	if err != nil {
		log.Errorf("error on UpdateRequestCredential:%s\n", err.Error())
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
		return
	}
	resp.Response(http.StatusOK, message.SUCCEED_CODE, "", nil)
	return
}

func (c *CredentialController) QueryCredential(ctx *gin.Context) {
	resp := common.Gin{C: ctx}
	data, isForward, err := common.ParseMessage(common.EnablePackage, ctx, c.packager, common.QueryCredentialType, c.msgSvr)
	if err != nil {
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
		return
	}
	if isForward {
		resp.Response(http.StatusOK, message.SUCCEED_CODE, "", nil)
		return
	}
	req, ok := data.(*message.QueryCredentialRequest)
	if !ok {
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, fmt.Errorf("data convert err").Error(), nil)
		return
	}
	rec, err := c.QueryCredentialFromStore(req.DId, req.Id)
	if err != nil {
		log.Errorf("error on QueryCredentialType:%s\n", err.Error())
		resp.Response(http.StatusOK, message.ERROR_CODE_INNER, err.Error(), nil)
		return
	}
	resp.Response(http.StatusOK, message.SUCCEED_CODE, "", &message.QueryCredentialResponse{
		CredentialsAttach: rec.CredentialsAttach,
		Formats:           rec.Formats,
	})
	return
}
