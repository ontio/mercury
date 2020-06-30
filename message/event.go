package message

type MessageType int

const (
	InvitationType MessageType = iota
	SendConnectionRequestType
	ConnectionRequestType
	ConnectionResponseType
	ConnectionACKType

	ProposalCredentialType
	OfferCredentialType
	RequestCredentialType
	IssueCredentialType
	CredentialACKType

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
