package vdri

const (
	Version                 = "1.0"
	InvitationSpec          = "spec/connections/" + Version + "/invitation"
	ConnectionRequestSpec   = "spec/connections/" + Version + "/request"
	ConnectionResponseSpec  = "spec/connections/" + Version + "/response"
	ConnectionACKSpec       = "spec/connections/" + Version + "/ack"
	BasicMsgSpec            = "spec/didcomm/" + Version + "/generalmessage"
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
