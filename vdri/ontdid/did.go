package ontdid

import (
	"encoding/json"
	"fmt"
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

func NewDID(ontSdk *sdk.OntologySdk, acc *sdk.Account) (string, error) {
	did, err := sdk.GenerateID()
	if err != nil {
		return "", err
	}
	err = RegisterDid(did, ontSdk, acc)
	if err != nil {
		return "", err
	}
	return did, nil
}

func RegisterDid(did string, ontSdk *sdk.OntologySdk, acc *sdk.Account) error {
	if ontSdk.Native == nil || ontSdk.Native.OntId == nil {
		return fmt.Errorf("ontsdk is nil")
	}
	_, err := ontSdk.Native.OntId.RegIDWithPublicKey(100, 100, acc, did, acc)
	if err != nil {
		return err
	}
	return nil
}

func ValidateDid(did string) bool {
	return sdk.VerifyID(did)
}

func GetDidDocByDid(did string, ontSdk *sdk.OntologySdk) (*Doc, error) {
	if ontSdk.Native == nil || ontSdk.Native.OntId == nil {
		return nil, fmt.Errorf("ontsdk is nil")
	}
	data, err := ontSdk.Native.OntId.GetDocumentJson(did)
	if err != nil {
		return nil, err
	}
	doc := &Doc{}
	err = json.Unmarshal(data, &doc)
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
	doc := &Doc{}
	err = json.Unmarshal(data, &doc)
	if err != nil {
		return nil, err
	}
	addrs := make([]string, 0)
	for _, endPoint := range doc.Service {
		addrs = append(addrs, endPoint.ServiceEndpoint)
	}
	return addrs, nil
}
