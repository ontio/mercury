package service

import (
	"fmt"
	"git.ont.io/ontid/otf/message"
	"time"
)

type ConnectionState int
type CredentialState int
type RequestCredentialState int

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
)

type InvitationRec struct {
	Invitation message.Invitation `json:"invitation"`
	State      ConnectionState    `json:"state"`
}

type ConnectionRequestRec struct {
	ConnReq message.ConnectionRequest `json:"conn_req"`
	State   ConnectionState           `json:"state"`
}

type ConnectionRec struct {
	OwnerDID    string `json:"owner_did"`
	Connections map[string]message.Connection
}

type RequestCredentialRec struct {
	RequesterDID      string                    `json:"requester_did"`
	RequestCredential message.RequestCredential `json:"request_credential"`
	State             RequestCredentialState    `json:"state"`
}

type CredentialRec struct {
	OwnerDID   string                  `json:"owner_did"`
	Credential message.IssueCredential `json:"credential"`
	Timestamp  time.Time               `json:"timestamp"`
}

type ServiceDoc struct {
	ServiceID       string `json:"service_id"`
	ServiceType     string `json:"service_type"`
	ServiceEndpoint string `json:"service_endpoint"`
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
		fmt.Printf("s.serviceid:%s\n", s.ServiceID)
		fmt.Printf("serviceID:%s\n", serviceID)
		if s.ServiceID == serviceID {
			return s.ServiceEndpoint, nil
		}
	}

	return "", fmt.Errorf("servicepoint not found")
}
