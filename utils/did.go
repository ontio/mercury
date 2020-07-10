package utils

import (
	"encoding/json"
	"fmt"
	"git.ont.io/ontid/otf/message"
	sdk "github.com/ontio/ontology-go-sdk"
)

type PublicKey struct {
	ID           string
	Type         string
	Controller   string
	PublicKeyHex string
}

type Authentication struct {
	Did       string
	PublicKey PublicKey
}

type ServiceDoc struct {
	ServiceID       string
	ServiceType     string
	ServiceEndpoint string
}

type Doc struct {
	Context        []string
	Id             string
	PublicKey      []PublicKey
	Authentication Authentication
	Controller     interface{}
	Recovery       interface{}
	Service        []ServiceDoc
	Attribute      interface{}
	Created        interface{}
	Updated        interface{}
	Proof          interface{}
}

func ValidateDid(did string) bool {
	return sdk.VerifyID(did)
}

func GetDidDocByDid(did string, ontSdk *sdk.OntologySdk) (*message.DIDDoc, error) {
	if ontSdk.Native == nil || ontSdk.Native.OntId == nil {
		return nil, fmt.Errorf("ontsdk is nil")
	}
	data, err := ontSdk.Native.OntId.GetDocumentJson(did)
	if err != nil {
		return nil, err
	}
	doc := &message.DIDDoc{}
	err = json.Unmarshal([]byte(string(data)), doc)
	if err != nil {
		return nil, err
	}
	return doc, nil
}

func GetServiceEndpointByDid(did string, ontSdk *sdk.OntologySdk) ([]string, error) {
	if ontSdk.Native == nil || ontSdk.Native.OntId == nil {
		return nil, fmt.Errorf("ontsdk is nil")
	}
	data, err := ontSdk.Native.OntId.GetDocumentJson(did)
	if err != nil {
		return nil, err
	}
	doc := &message.DIDDoc{}
	err = json.Unmarshal([]byte(string(data)), &doc)
	if err != nil {
		return nil, err
	}
	addrs := make([]string, 0)
	for _, endPoint := range doc.Service {
		addrs = append(addrs, endPoint.ServiceEndpoint)
	}
	return addrs, nil
}

func GetPubKeyByDid(did string) (string, error) {
	return "", nil
}
