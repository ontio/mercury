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

package vdri

const (
	Version                 = "1.0"
	InvitationSpec          = "spec/connections/" + Version + "/invitation"
	ConnectionRequestSpec   = "spec/connections/" + Version + "/request"
	ConnectionResponseSpec  = "spec/connections/" + Version + "/response"
	ConnectionACKSpec       = "spec/connections/" + Version + "/ack"
	BasicMsgSpec            = "spec/didcomm/" + Version + "/basicmessage"
	ProposalCredentialSpec  = "spec/issue-credential/" + Version + "/propose-credential"
	OfferCredentialSpec     = "spec/issue-credential/" + Version + "/offer-credential"
	RequestCredentialSpec   = "spec/issue-credential/" + Version + "/request-credential"
	IssueCredentialSpec     = "spec/issue-credential/" + Version + "/issue-credential"
	CredentialACKSpec       = "spec/issue-credential/" + Version + "/ack"
	RequestPresentationSpec = "spec/present-proof/" + Version + "/request-presentation"
	PresentationProofSpec   = "spec/present-proof/" + Version + "/presentation"
	PresentationACKSpec     = "spec/present-proof/" + Version + "/ack"
)

type DidDoc interface {
	GetServicePoint(serviceId string) (string, error)
	GetServiceEndpointByDid(did string, sdk interface{}) ([]string, error)
	GetDidDocByDid(did string, sdk interface{}) (interface{}, error)
}
