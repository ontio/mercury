package common

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
	DeleteCredentialType
	QueryCredentialType

	SendRequestPresentationType
	RequestPresentationType
	PresentationType
	PresentationAckType
	DeletePresentationType
	QueryPresentationType

	SendBasicMsgType
	ReceiveBasicMsgType
	QueryBasicMessageType
)

type Message struct {
	MessageType `json:"type"`
	Content     interface{} `json:"content"`
}

const (
	InviteApi                    = "/api/v1/invitation"
	ConnectRequestApi            = "/api/v1/connectionrequest"
	ConnectResponseApi           = "/api/v1/connectionresponse"
	ConnectAckApi                = "/api/v1/connectionack"
	SendDisconnectApi            = "/api/v1/senddisconnect"
	DisconnectApi                = "/api/v1/disconnect"
	SendProposalCredentialReqApi = "/api/v1/sendproposalcredential"
	ProposalCredentialReqApi     = "/api/v1/proposalcredential"
	OfferCredentialApi           = "/api/v1/offercredential"
	SendRequestCredentialApi     = "/api/v1/sendrequestcredential"
	RequestCredentialApi         = "/api/v1/requestcredential"
	IssueCredentialApi           = "/api/v1/issuecredentail"
	CredentialAckApi             = "/api/v1/credentialack"
	DeleteCredetialApi           = "/api/v1/deletecredential"
	QueryCredentialApi           = "/api/v1/querycredential"
	SendRequestPresentationApi   = "/api/v1/sendrequestpresentation"
	RequestPresentationApi       = "/api/v1/requestpresentation"
	PresentationProofApi         = "/api/v1/presentproof"
	PresentationAckApi           = "/api/v1/presentationack"
	QueryPresentationApi         = "/api/v1/querypresentation"
	DeletePresentationApi        = "/api/v1/deletepresentation"
	SendBasicMsgApi              = "/api/v1/sendbasicmsg"
	ReceiveBasicMsgApi           = "/api/v1/receivebasicmsg"
	QueryBasicMsgApi             = "/api/v1/querybasicmsg"
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

func TransferForwardMsgType(msgType MessageType) MessageType {
	switch msgType {
	case SendBasicMsgType:
		return ReceiveBasicMsgType
	case SendRequestCredentialType:
		return RequestCredentialType
	case SendProposalCredentialType:
		return ProposalCredentialType
	case SendDisconnectType:
		return DisconnectType
	case SendRequestPresentationType:
		return RequestPresentationType
	default:
		return msgType
	}
}
