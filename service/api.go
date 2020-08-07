/*
 * Copyright (C) 2018 The ontology Authors
 * This file is part of The ontology library.
 *
 * The ontology is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The ontology is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Lesser General Public License for more details.
 *
 * You should have received a copy of the GNU Lesser General Public License
 * along with The ontology.  If not, see <http://www.gnu.org/licenses/>.
 */

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
