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

func (self *Invitation) GetConnection() *Connection {
	return nil
}

type RequestInf interface {
	GetConnection() *Connection
}

type ConnectionRequest struct {
	Type         string     `json:"@type,omitempty"`
	Id           string     `json:"@id,omitempty"`
	Label        string     `json:"label,omitempty"`
	Connection   Connection `json:"connection,omitempty"`
	InvitationId string     `json:"invitation_id"`
}

func (self *ConnectionRequest) GetConnection() *Connection {
	return &self.Connection
}

type Connection struct {
	MyDid       string   `json:"my_did,omitempty"`
	MyRouter    []string `json:"my_router"`
	TheirDid    string   `json:"their_did"`
	TheirRouter []string `json:"their_router"`
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

func (self *ConnectionResponse) GetConnection() *Connection {
	return &self.Connection
}

type ConnectionACK struct {
	Type       string     `json:"@type,omitempty"`
	Id         string     `json:"@id,omitempty"`
	Thread     Thread     `json:"~thread,omitempty"`
	Status     string     `json:"status,omitempty"`
	Connection Connection `json:"connection,omitempty"`
}

func (self *ConnectionACK) GetConnection() *Connection {
	return &self.Connection
}

type DisconnectRequest struct {
	Type       string     `json:"@type,omitempty"`
	Id         string     `json:"@id,omitempty"`
	Connection Connection `json:"connection,omitempty"`
}

func (self *DisconnectRequest) GetConnection() *Connection {
	return &self.Connection
}

type CredentialACK struct {
	Type       string     `json:"@type,omitempty"`
	Id         string     `json:"@id,omitempty"`
	Thread     Thread     `json:"~thread,omitempty"`
	Status     string     `json:"status,omitempty"`
	Connection Connection `json:"connection,omitempty"`
}

func (self *CredentialACK) GetConnection() *Connection {
	return &self.Connection
}

type PresentationACK struct {
	Type       string     `json:"@type,omitempty"`
	Id         string     `json:"@id,omitempty"`
	Thread     Thread     `json:"~thread,omitempty"`
	Status     string     `json:"status,omitempty"`
	Connection Connection `json:"connection,omitempty"`
}

func (self *PresentationACK) GetConnection() *Connection {
	return &self.Connection
}

//issue credential
type ProposalCredential struct {
	Type               string            `json:"@type,omitempty"`
	Id                 string            `json:"@id,omitempty"`
	Comment            string            `json:"comment,omitempty"`
	CredentialProposal CredentialPreview `json:"credential_proposal,omitempty"`
	Connection         Connection        `json:"connection,omitempty"`
}

func (self *ProposalCredential) GetConnection() *Connection {
	return &self.Connection
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

func (self *OfferCredential) GetConnection() *Connection {
	return &self.Connection
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

func (self *RequestCredential) GetConnection() *Connection {
	return &self.Connection
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

func (self *IssueCredential) GetConnection() *Connection {
	return &self.Connection
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

func (self *RequestPresentation) GetConnection() *Connection {
	return &self.Connection
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

func (self *Presentation) GetConnection() *Connection {
	return &self.Connection
}

type BasicMessage struct {
	Type       string     `json:"@type"`
	Id         string     `json:"@id"`
	SendTime   time.Time  `json:"send_time"`
	Content    string     `json:"content"`
	I10n       I10n       `json:"~I10n"`
	Connection Connection `json:"connection,omitempty"`
}

func (self *BasicMessage) GetConnection() *Connection {
	return &self.Connection
}

type I10n struct {
	Locale string `json:"locale"`
}

type DeleteCredentialRequest struct {
	DId string `json:"did"`
	Id  string `json:"id"`
}

func (self *DeleteCredentialRequest) GetConnection() *Connection {
	return nil
}

type DeletePresentationRequest struct {
	DId string `json:"did"`
	Id  string `json:"id"`
}

func (self *DeletePresentationRequest) GetConnection() *Connection {
	return nil
}

type QueryCredentialRequest struct {
	DId string `json:"did"`
	Id  string `json:"id"`
}

func (self *QueryCredentialRequest) GetConnection() *Connection {
	return nil
}

type QueryCredentialResponse struct {
	Formats           []Format     `json:"formats,omitempty"`
	CredentialsAttach []Attachment `json:"credentials~attach,omitempty"`
}

func (self *QueryCredentialResponse) GetConnection() *Connection {
	return nil
}

type QueryPresentationRequest struct {
	DId string `json:"did"`
	Id  string `json:"id"`
}

func (self *QueryPresentationRequest) GetConnection() *Connection {
	return nil
}

type QueryPresentationResponse struct {
	Formats            []Format     `json:"formats,omitempty"`
	PresentationAttach []Attachment `json:"presentations~attach,omitempty"`
}

func (self *QueryPresentationResponse) GetConnection() *Connection {
	return nil
}

type QueryGeneralMessageRequest struct {
	DID             string `json:"did"`
	Latest          bool   `json:"latest"`
	RemoveAfterRead bool   `json:"remove_after_read"`
}

func (self *QueryGeneralMessageRequest) GetConnection() *Connection {
	return nil
}

type ForwardMessageRequest struct {
	MsgType    int        `json:"msg_type"`
	Data       []byte     `json:"data"`
	Connection Connection `json:"connection"`
}

func (self *ForwardMessageRequest) GetConnection() *Connection {
	return &self.Connection
}
