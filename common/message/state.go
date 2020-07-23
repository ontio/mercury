package message

import (
	"fmt"
	"time"
)

type ConnectionState int
type CredentialState int
type RequestCredentialState int
type RequestPresentationState int

const (
	InvitationInit ConnectionState = iota
	InvitationUsed
	ConnectionRequestSent
	ConnectionRequestReceived
	ConnectionResponseReceived
	ConnectionACKReceived

	RequestCredentialReceived RequestCredentialState = iota
	RequestCredentialResolved

	CredentialIssued CredentialState = iota
	CredentialReceive

	RequestPresentationReceived RequestPresentationState = iota
	RequestPresentationResolved
)

type BasicMsgRec struct {
	Msglist []BasicMessage `json:"msglist"`
}

type InvitationRec struct {
	Invitation Invitation      `json:"invitation"`
	State      ConnectionState `json:"state"`
}

type ConnectionRequestRec struct {
	ConnReq ConnectionRequest `json:"conn_req"`
	State   ConnectionState   `json:"state"`
}

type ConnectionRec struct {
	OwnerDID    string                `json:"owner_did"`
	Connections map[string]Connection `json:"connections"`
}

type RequestCredentialRec struct {
	RequesterDID      string                 `json:"requester_did"`
	RequestCredential RequestCredential      `json:"request_credential"`
	State             RequestCredentialState `json:"state"`
}

type CredentialRec struct {
	OwnerDID   string          `json:"owner_did"`
	Credential IssueCredential `json:"credential"`
	Timestamp  time.Time       `json:"timestamp"`
}

type ServiceDoc struct {
	ServiceID       string `json:"id"`
	ServiceType     string `json:"type"`
	ServiceEndpoint string `json:"serviceEndpint"` //todo fix this typo with ontology update
}

type RequestPresentationRec struct {
	RequesterDID       string                   `json:"requester_did"`
	RerquestPrentation RequestPresentation      `json:"rerquest_prentation"`
	State              RequestPresentationState `json:"state"`
}

type PresentationRec struct {
	OwnerDID     string       `json:"owner_did"`
	Presentation Presentation `json:"presentation"`
	Timestamp    time.Time    `json:"timestamp"`
}

type DIDDoc struct {
	Context        []string     `json:"@context"`
	Id             string       `json:"id"`
	PublicKey      interface{}  `json:"publicKey"`
	Authentication interface{}  `json:"authentication"`
	Controller     interface{}  `json:"controller"`
	Recovery       interface{}  `json:"recovery"`
	Service        []ServiceDoc `json:"service"`
	Attribute      interface{}  `json:"attribute"`
	Created        interface{}  `json:"created"`
	Updated        interface{}  `json:"updated"`
	Proof          interface{}  `json:"proof"`
}

func (d DIDDoc) GetServicePoint(serviceID string) (string, error) {
	for _, s := range d.Service {
		if s.ServiceID == serviceID {
			return s.ServiceEndpoint, nil
		}
	}
	return "", fmt.Errorf("servicepoint not found")
}
