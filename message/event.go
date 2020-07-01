package message

type MessageType int

const (
	InvitationType MessageType = iota
	SendConnectionRequestType
	ConnectionRequestType
	ConnectionResponseType
	ConnectionACKType

	SendProposalCredentialType
	ProposalCredentialType
	OfferCredentialType
	SendRequestCredentialType
	RequestCredentialType
	IssueCredentialType
	CredentialACKType

	SendPresentationType
	RequestPresentationType
	PresentationType
	PresentationACKType

	SendGeneralMsgType
	ReceiveGeneralMsgType
)

type Message struct {
	MessageType `json:"message_type"`
	Content     interface{} `json:"content"`
	JsonBytes   []byte      `json:"json_bytes,,omitempty"`
}
