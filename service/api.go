package service

import "github.com/gin-gonic/gin"

type SystemApiServicer interface {
	Invitation(ctx *gin.Context)
	ConnectionRequest(ctx *gin.Context)
	ConnectionResponse(ctx *gin.Context)
	ConnectionAck(ctx *gin.Context)
	SendDisConnect(ctx *gin.Context)
	Disconnect(ctx *gin.Context)
	SendBasicMsg(ctx *gin.Context)
	ReceiveBasicMsg(ctx *gin.Context)
	QueryBasicMsg(ctx *gin.Context)
}

type CredentialApiServicer interface {
	SendProposalCredential(ctx *gin.Context)
	ProposalCredential(ctx *gin.Context)
	OfferCredential(ctx *gin.Context)
	SendRequestCredential(ctx *gin.Context)
	RequestCredential(ctx *gin.Context)
	IssueCredential(ctx *gin.Context)
	CredentialAck(ctx *gin.Context)
}

type PresentationApiServicer interface {
	SendRequestPresentation(ctx *gin.Context)
	RequestPresentation(ctx *gin.Context)
	PresentProof(ctx *gin.Context)
	PresentationAck(ctx *gin.Context)
	QueryPresentation(ctx *gin.Context)
}
