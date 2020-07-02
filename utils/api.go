package utils

import "git.ont.io/ontid/otf/message"

const (
	Group_Api_V1                  = "/api/v1"
	Invite_Api                    = "/invitation"
	SendConnectionReq_Api         = "/sendconnectionreq"
	ConnectRequest_Api            = "/connectionrequest"
	ConnectResponse_Api           = "/connectionresponse"
	ConnectAck_Api                = "/connectionack"
	SendProposalCredentialReq_Api = "/sendproposalcredential"
	ProposalCredentialReq_Api     = "/proposalcredential"
	OfferCredential_Api           = "/offercredential"
	SendRequestCredential_Api     = "/sendrequestcredential"
	RequestCredential_Api         = "/requestcredential"
	IssueCredential_Api           = "/issuecredentail"
	CredentialAckInfo_Api         = "/credentialack"
	SendRequestPresentation_Api   = "/sendRequestPresentation"
	RequestPresentation_Api       = "/requestpresentation"
	Presentation_Api              = "/presentproof"
	PresentationAckInfo           = "/presentationack"
	SendGeneralMsg                = "/sendgeneralmsg"
	ReceiveGeneralMsg             = "/receivegeneralmsg"
)

func GetApiName(msgType message.MessageType) string {
	switch msgType {
	case message.InvitationType:
		return Group_Api_V1 + Invite_Api
	case message.SendConnectionRequestType:
		return Group_Api_V1 + SendConnectionReq_Api
	case message.ConnectionRequestType:
		return Group_Api_V1 + ConnectRequest_Api
	case message.ConnectionResponseType:
		return Group_Api_V1 + ConnectResponse_Api
	case message.ConnectionACKType:
		return Group_Api_V1 + ConnectAck_Api
	case message.ProposalCredentialType:
		return Group_Api_V1 + ProposalCredentialReq_Api
	case message.OfferCredentialType:
		return Group_Api_V1 + OfferCredential_Api
	case message.RequestCredentialType:
		return Group_Api_V1 + RequestCredential_Api
	case message.IssueCredentialType:
		return Group_Api_V1 + IssueCredential_Api
	case message.CredentialACKType:
		return Group_Api_V1 + CredentialAckInfo_Api
	case message.RequestPresentationType:
		return Group_Api_V1 + RequestPresentation_Api
	case message.PresentationType:
		return Group_Api_V1 + Presentation_Api
	case message.PresentationACKType:
		return Group_Api_V1 + PresentationAckInfo
	case message.SendGeneralMsgType:
		return Group_Api_V1 + ReceiveGeneralMsg
	default:
		return Group_Api_V1
	}
	return Group_Api_V1
}
