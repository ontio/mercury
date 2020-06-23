package rest

type Invitation struct {
	Type            string   `json:"@type"`
	Id              string   `json:"@id"`
	Lable           string   `json:"lable"`
	Did             string   `json:"did"`
	RecipientKeys   []string `json:"recipientKeys"`
	ServiceEndpoint string   `json:"serviceEndpoint"`
	routingKeys     string   `json:"routingKeys"`
}

type ConnectionRequest struct {
	Type       string     `json:"@type"`
	Id         string     `json:"@id"`
	Lable      string     `json:"lable"`
	Connection Connection `json:"connection"`
}

type Connection struct {
	Did    string `json:"did"`
	DIDDoc DIDDoc `json:"diddoc"`
}

type DIDDoc struct {
	Context   []string    `json:"@context"`
	Id        string      `json:"id"`
	PublicKey []PublicKey `json:"publicKey"`
}

type PublicKey struct {
	Type            string           `json:"type"`
	Id              string           `json:"id"`
	Controller      string           `json:"controller"`
	PublicKeyBase58 string           `json:"publicKeyBase58"`
	Authentication  []Authentication `json:"authentication"`
	Service         Service          `json:"service"`
}

type Authentication struct {
	Type      string `json:"type"`
	PublicKey string `json:"publicKey"`
}

type Service struct {
	Id              string   `json:"id"`
	Type            string   `json:"type"`
	Priority        int      `json:"priority"`
	RecipientKeys   []string `json:"recipientKeys"`
	ServiceEndpoint string   `json:"serviceEndpoint"`
}

type ConnectResponse struct {
	Type   string `json:"@type"`
	Id     string `json:"@id"`
	Thread Thread `json:"~thread"`
}

type Thread struct {
	Thid       string     `json:"thid"`
	Connection Connection `json:"connection"`
}

type ConnectionACK struct {
	Type   string `json:"@type"`
	Id     string `json:"@id"`
	Thread Thread `json:"~thread"`
}

//========issue credential
type ProposalCredential struct {
	Type               string            `json:"@type"`
	Id                 string            `json:"@id"`
	Comment            string            `json:"comment"`
	CredentialProposal CredentialPreview `json:"credential_proposal"`
	SchemaIssuerDid    string            `json:"schema_issuer_did"`
	SchemaId           string            `json:"schema_id"`
	SchemaVersion      string            `json:"schema_version"`
	CredDefId          string            `json:"cred_def_id"`
	IssuerDid          string            `json:"issuer_did"`
}

type OfferCredential struct {
	Type              string            `json:"@type"`
	Id                string            `json:"@id"`
	Comment           string            `json:"comment"`
	CredentialPreview CredentialPreview `json:"credential_preview"`
	OffersAttach      []Attachment      `json:"offers_attach"`
}

type CredentialPreview struct {
	Type       string       `json:"@type"`
	Attributre []Attributre `json:"attributre"`
}

type Attributre struct {
	Name     string `json:"name"`
	MimeType string `json:"mime_type"`
	Value    string `json:"value"`
	CredDefId          string            `json:"cred_def_id"`
	referent string `json:"referent"`
}

type Attachment struct {
	Id       string `json:"@id"`
	MimeType string `json:"mime_type"`
	Data     Data   `json:"data"`
}

type Data struct {
	Base64 string `json:"base64"`
}

type RequestCredential struct {
	Type           string       `json:"@type"`
	Id             string       `json:"@id"`
	Comment        string       `json:"comment"`
	RequestsAttach []Attachment `json:"requests_attach"`
}

type CredentialACK struct {
	Type   string `json:"@type"`
	Id     string `json:"@id"`
	Thread Thread `json:"~thread"`
}

//present proof
type ProposPresentation struct {
	Type    string `json:"@type"`
	Id      string `json:"@id"`
	Comment string `json:"comment"`
	PresentationProposal PresentationPreview `json:"presentation_proposal"`
}

type PresentationPreview struct {
	Type    string `json:"@type"`
	Attributes []Attributre `json:"attributes"`
	Predicates Predicate `json:"predicates"`
}

type Predicate struct {
	Name string `json:"name"`
	CredDefId          string            `json:"cred_def_id"`
	Predicate string `json:"predicate"`
	Threshold string `json:"threshold"`
}

type RequestPresentation struct {
	Type    string `json:"@type"`
	Id      string `json:"@id"`
	Comment string `json:"comment"`
	RequestPresentationAttach []Attachment `json:"request_presentation_attach"`
}

type Presentation struct {
	Type    string `json:"@type"`
	Id      string `json:"@id"`
	Comment string `json:"comment"`
	PresentationAttach []Attachment `json:"presentation_attach"`
}

type PresentationACK struct {
	Type   string `json:"@type"`
	Id     string `json:"@id"`
	Thread Thread `json:"~thread"`
}