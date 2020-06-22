package message

type MessageType int

const (
	Invitation MessageType = iota
	ConnectionRequest
	ConnectionResponse
	ConnectionACK

	ProposalCredential
	OfferCredential
	RequestCredential
	IssueCredential
	CredentialACK

	RequestPresentation
	Presentation
	PresentationACK
)

type Message struct {
	MessageType
	Content map[string]interface{}
}
