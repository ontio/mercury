package common

import (
	"fmt"
	"git.ont.io/ontid/otf/common/message"
)

type MessageType int

const (
	InvitationType MessageType = iota
	ConnectionRequestType
	ConnectionResponseType
	ConnectionAckType
	SendDisconnectType
	DisconnectType

	SendProposalCredentialType
	ProposalCredentialType
	OfferCredentialType
	SendRequestCredentialType
	RequestCredentialType
	IssueCredentialType
	CredentialAckType

	SendRequestPresentationType
	RequestPresentationType
	PresentationType
	PresentationAckType
	SendBasicMsgType
	ReceiveBasicMsgType
	QueryBasicMessageType

	QueryCredentialType
	QueryPresentationType
)

type Message struct {
	MessageType `json:"type"`
	Content     interface{} `json:"content"`
}

const (
	InviteApi          = "/api/v1/invitation"
	ConnectRequestApi  = "/api/v1/connectionrequest"
	ConnectResponseApi = "/api/v1/connectionresponse"
	ConnectAckApi      = "/api/v1/connectionack"
	SendDisconnectApi  = "/api/v1/senddisconnect"
	DisconnectApi      = "/api/v1/disconnect"

	SendProposalCredentialReqApi = "/api/v1/sendproposalcredential"
	ProposalCredentialReqApi     = "/api/v1/proposalcredential"
	OfferCredentialApi           = "/api/v1/offercredential"
	SendRequestCredentialApi     = "/api/v1/sendrequestcredential"
	RequestCredentialApi         = "/api/v1/requestcredential"
	IssueCredentialApi           = "/api/v1/issuecredentail"
	CredentialAckApi             = "/api/v1/credentialack"

	SendRequestPresentationApi = "/api/v1/sendrequestpresentation"
	RequestPresentationApi     = "/api/v1/requestpresentation"
	PresentationProofApi       = "/api/v1/presentproof"
	PresentationAckApi         = "/api/v1/presentationack"
	SendBasicMsgApi            = "/api/v1/sendbasicmsg"
	ReceiveBasicMsgApi         = "/api/v1/receivebasicmsg"
	QueryBasicMsgApi           = "/api/v1/querybasicmsg"
	QueryCredentialApi         = "/api/v1/querycredential"
	QueryPresentationApi       = "/api/v1/querypresentation"
)

func GetApiName(msgType MessageType) string {
	switch msgType {
	case InvitationType:
		return InviteApi
	case ConnectionRequestType:
		return ConnectRequestApi
	case ConnectionResponseType:
		return ConnectResponseApi
	case ConnectionAckType:
		return ConnectAckApi
	case DisconnectType:
		return DisconnectApi
	case ProposalCredentialType:
		return ProposalCredentialReqApi
	case OfferCredentialType:
		return OfferCredentialApi
	case RequestCredentialType:
		return RequestCredentialApi
	case IssueCredentialType:
		return IssueCredentialApi
	case CredentialAckType:
		return CredentialAckApi
	case RequestPresentationType:
		return RequestPresentationApi
	case PresentationType:
		return PresentationProofApi
	case PresentationAckType:
		return PresentationAckApi
	case SendBasicMsgType:
		return SendBasicMsgApi
	case ReceiveBasicMsgType:
		return ReceiveBasicMsgApi
	case QueryBasicMessageType:
		return QueryBasicMsgApi
	case QueryCredentialType:
		return QueryCredentialApi
	case QueryPresentationType:
		return QueryPresentationApi
	default:
		return ""
	}
}

func getMsgObjectByType(messageType MessageType) (interface{}, error) {
	switch messageType {
	case InvitationType:
		return &message.Invitation{}, nil
	case ConnectionRequestType:
		return &message.ConnectionRequest{}, nil
	case ConnectionResponseType:
		return &message.ConnectionResponse{}, nil
	case ConnectionAckType:
		return &message.ConnectionACK{}, nil
	case DisconnectType, SendDisconnectType:
		return &message.DisconnectRequest{}, nil
	case SendProposalCredentialType:
		return &message.ProposalCredential{}, nil
	case OfferCredentialType:
		return &message.OfferCredential{}, nil
	case ProposalCredentialType:
		return &message.ProposalCredential{}, nil
	case SendRequestCredentialType:
		return &message.RequestCredential{}, nil
	case RequestCredentialType:
		return &message.RequestCredential{}, nil
	case IssueCredentialType:
		return &message.IssueCredential{}, nil
	case CredentialAckType:
		return &message.CredentialACK{}, nil
	case RequestPresentationType:
		return &message.RequestPresentation{}, nil
	case SendRequestPresentationType:
		return &message.RequestPresentation{}, nil
	case PresentationType:
		return &message.Presentation{}, nil
	case PresentationAckType:
		return &message.PresentationACK{}, nil
	case SendBasicMsgType:
		return &message.BasicMessage{}, nil
	case QueryBasicMessageType:
		return &message.QueryGeneralMessageRequest{}, nil
	case QueryCredentialType:
		return &message.QueryCredentialRequest{}, nil
	case QueryPresentationType:
		return &message.QueryPresentationRequest{}, nil
	default:
		return nil, fmt.Errorf("msg type err:%v", messageType)
	}
}
