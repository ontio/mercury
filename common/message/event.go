//message for rest api
package message

type MessageType int

const (
	InvitationType MessageType = iota
	ConnectionRequestType
	ConnectionResponseType
	ConnectionACKType
	SendDisconnectType
	DisconnectType

	SendProposalCredentialType
	ProposalCredentialType
	OfferCredentialType
	SendRequestCredentialType
	RequestCredentialType
	IssueCredentialType
	CredentialACKType

	SendRequestPresentationType
	RequestPresentationType
	PresentationType
	PresentationACKType

	SendGeneralMsgType
	ReceiveGeneralMsgType
	QueryGeneralMessageType

	QueryCredentialType
	QueryPresentationType
)

type Message struct {
	MessageType `json:"message_type"`
	Content     interface{} `json:"content"`
}