package message

type MessageType int

const (
	InvitationType MessageType = iota
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
)

type Message struct {
	MessageType
	Content map[string]interface{}
}
