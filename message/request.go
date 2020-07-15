package message

import "time"

type Invitation struct {
	Type   string   `json:"@type,omitempty"`
	Id     string   `json:"@id,omitempty"`
	Label  string   `json:"label,omitempty"`
	Did    string   `json:"did,omitempty"`
	Router []string `json:"router"`
	//ServiceId string   `json:"service_id,omitempty"`
}

type ConnectionRequest struct {
	Type         string     `json:"@type,omitempty"`
	Id           string     `json:"@id,omitempty"`
	Label        string     `json:"label,omitempty"`
	Connection   Connection `json:"connection,omitempty"`
	InvitationId string     `json:"invitation_id"`
}

type Connection struct {
	MyDid       string   `json:"my_did,omitempty"`
	MyRouter    []string `json:"my_router"`
	TheirDid    string   `json:"their_did"`
	TheirRouter []string `json:"their_router"`
	//MyServiceId    string   `json:"my_service_id,omitempty"`
	//TheirServiceId string   `json:"their_service_id"`
}

// Thread thread data
type Thread struct {
	ID             string         `json:"thid,omitempty"`
	PID            string         `json:"pthid,omitempty"`
	SenderOrder    int            `json:"sender_order,omitempty"`
	ReceivedOrders map[string]int `json:"received_orders,omitempty"`
}

type ConnectionResponse struct {
	Type       string     `json:"@type,omitempty"`
	Id         string     `json:"@id,omitempty"`
	Thread     Thread     `json:"~thread,omitempty"`
	Connection Connection `json:"connection,omitempty"`
}

type ConnectionACK struct {
	Type       string     `json:"@type,omitempty"`
	Id         string     `json:"@id,omitempty"`
	Thread     Thread     `json:"~thread,omitempty"`
	Status     string     `json:"status,omitempty"`
	Connection Connection `json:"connection,omitempty"`
}

type DisconnectRequest struct {
	Type       string     `json:"@type,omitempty"`
	Id         string     `json:"@id,omitempty"`
	Connection Connection `json:"connection,omitempty"`
}

type CredentialACK struct {
	Type       string     `json:"@type,omitempty"`
	Id         string     `json:"@id,omitempty"`
	Thread     Thread     `json:"~thread,omitempty"`
	Status     string     `json:"status,omitempty"`
	Connection Connection `json:"connection,omitempty"`
}
type PresentationACK struct {
	Type       string     `json:"@type,omitempty"`
	Id         string     `json:"@id,omitempty"`
	Thread     Thread     `json:"~thread,omitempty"`
	Status     string     `json:"status,omitempty"`
	Connection Connection `json:"connection,omitempty"`
}

//issue credential
type ProposalCredential struct {
	Type               string            `json:"@type,omitempty"`
	Id                 string            `json:"@id,omitempty"`
	Comment            string            `json:"comment,omitempty"`
	CredentialProposal CredentialPreview `json:"credential_proposal,omitempty"`
	Connection         Connection        `json:"connection,omitempty"`
}

type OfferCredential struct {
	Type              string            `json:"@type,omitempty"`
	Id                string            `json:"@id,omitempty"`
	Comment           string            `json:"comment,omitempty"`
	CredentialPreview CredentialPreview `json:"credential_preview,omitempty"`
	OffersAttach      []Attachment      `json:"offers_attach,omitempty"`
	Connection        Connection        `json:"connection,omitempty"`
	Thread            Thread            `json:"~thread,omitempty"`
}

type CredentialPreview struct {
	Type       string       `json:"@type,omitempty"`
	Attributre []Attributre `json:"attributre,omitempty"`
}

type Attributre struct {
	Name      string `json:"name,omitempty"`
	MimeType  string `json:"mime_type,omitempty"`
	Value     string `json:"value,omitempty"`
	CredDefId string `json:"cred_def_id,omitempty"`
	referent  string `json:"referent,omitempty"`
}

type Attachment struct {
	Id          string    `json:"@id,omitempty"`
	Description string    `json:"description,omitempty"`
	FileName    string    `json:"filename,omitempty"`
	MimeType    string    `json:"mime_type,omitempty"`
	LastModTime time.Time `json:"lastmod_time,omitempty"`
	ByteCount   int64     `json:"byte_count,omitempty"`
	Data        Data      `json:"data,omitempty"`
}

type Data struct {
	Sha256 string      `json:"sha256,omitempty"`
	Links  []string    `json:"links,omitempty"`
	Base64 string      `json:"base64,omitempty"`
	JSON   interface{} `json:"json,omitempty"`
}

type Format struct {
	AttachID string `json:"attach_id,omitempty"`
	Format   string `json:"format,omitempty"`
}

type RequestCredential struct {
	Type           string       `json:"@type"`
	Id             string       `json:"@id"`
	Comment        string       `json:"comment"`
	Formats        []Format     `json:"formats,omitempty"`
	RequestsAttach []Attachment `json:"requests_attach"`
	Connection     Connection   `json:"connection,omitempty"`
}

type IssueCredential struct {
	Type              string       `json:"@type,omitempty"`
	Id                string       `json:"@id"`
	Comment           string       `json:"comment,omitempty"`
	Formats           []Format     `json:"formats,omitempty"`
	CredentialsAttach []Attachment `json:"credentials~attach,omitempty"`
	Connection        Connection   `json:"connection,omitempty"`
	Thread            Thread       `json:"~thread,omitempty"`
}

//present proof
type ProposePresentation struct {
	Type          string       `json:"@type,omitempty"`
	Id            string       `json:"@id,omitempty"`
	Comment       string       `json:"comment,omitempty"`
	Formats       []Format     `json:"formats,omitempty"`
	ProposeAttach []Attachment `json:"propose~attach,omitempty"`
}

type RequestPresentation struct {
	Type                      string       `json:"@type,omitempty"`
	Id                        string       `json:"@id,omitempty"`
	Comment                   string       `json:"comment,omitempty"`
	Formats                   []Format     `json:"formats,omitempty"`
	RequestPresentationAttach []Attachment `json:"request_presentation_attach,omitempty"`
	Connection                Connection   `json:"connection,omitempty"`
}

type Presentation struct {
	Type               string       `json:"@type,omitempty"`
	Id                 string       `json:"@id,omitempty"`
	Comment            string       `json:"comment,omitempty"`
	Formats            []Format     `json:"formats,omitempty"`
	PresentationAttach []Attachment `json:"presentations~attach,omitempty"`
	Connection         Connection   `json:"connection,omitempty"`
	Thread             Thread       `json:"~thread,omitempty"`
}

type BasicMessage struct {
	Type       string     `json:"@type"`
	Id         string     `json:"@id"`
	SendTime   time.Time  `json:"send_time"`
	Content    string     `json:"content"`
	I10n       I10n       `json:"~I10n"`
	Connection Connection `json:"connection,omitempty"`
}

type I10n struct {
	Locale string `json:"locale"`
}

type QueryCredentialRequest struct {
	DId string `json:"did"`
	Id  string `json:"id"`
}

type QueryCredentialResponse struct {
	Formats           []Format     `json:"formats,omitempty"`
	CredentialsAttach []Attachment `json:"credentials~attach,omitempty"`
}

type QueryPresentationRequest struct {
	DId string `json:"did"`
	Id  string `json:"id"`
}

type QueryPresentationResponse struct {
	Formats            []Format     `json:"formats,omitempty"`
	PresentationAttach []Attachment `json:"presentations~attach,omitempty"`
}

type QueryGeneralMessageRequest struct {
	DID             string `json:"did"`
	Latest          bool   `json:"latest"`
	RemoveAfterRead bool   `json:"remove_after_read"`
}
