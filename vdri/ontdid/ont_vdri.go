package ontdid

import (
	"encoding/json"
	"git.ont.io/ontid/otf/message"
	"git.ont.io/ontid/otf/service"
	"git.ont.io/ontid/otf/vdri"
	"github.com/google/uuid"
	sdk "github.com/ontio/ontology-go-sdk"
	"time"
)

type SampleSubject struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

//todo use sdk query from smart contract
type OntVDRI struct {
	ontSdk *sdk.OntologySdk
	acct   *sdk.Account
	did    string
}

func NewOntVDRI(ontsdk *sdk.OntologySdk, acct *sdk.Account, did string) *OntVDRI {
	return &OntVDRI{
		ontSdk: ontsdk,
		acct:   acct,
		did:    did,
	}
}

func (ontVdri *OntVDRI) OfferCredential(req *message.ProposalCredential) (*message.OfferCredential, error) {
	return nil, nil
}
func (ontVdri *OntVDRI) IssueCredential(req *message.RequestCredential) (*message.IssueCredential, error) {

	//for test
	subs := make([]*SampleSubject, 0)
	for _, attach := range req.RequestsAttach {
		s := attach.Data.JSON.(string)
		sample := new(SampleSubject)
		err := json.Unmarshal([]byte(s), sample)
		if err != nil {
			return nil, err
		}
		subs = append(subs, sample)
	}
	contexts := []string{"context1", "context2"}
	types := []string{"SampleCredential"}
	expirationDate := time.Now().UTC().Unix() + 86400

	vc, err := ontVdri.ontSdk.Credential.CreateJWTCredential(contexts, types, subs, ontVdri.did, expirationDate, "", nil, ontVdri.acct)
	if err != nil {
		return nil, err
	}

	//fixme
	credential := &message.IssueCredential{
		Type:    vdri.IssueCredentialSpec,
		Id:      uuid.New().String(),
		Comment: "ontdid issueCredential",
		Formats: []message.Format{message.Format{
			AttachID: "1",
			Format:   "base64",
		}},
		CredentialsAttach: []message.Attachment{message.Attachment{
			Id:          "1",
			LastModTime: time.Now(),
			Data: message.Data{
				Base64: vc,
			},
		}},
		Connection: service.ReverseConnection(req.Connection),
		Thread: message.Thread{
			ID: req.Id,
		},
	}
	//todo do we need to commit credential to blockchain?

	return credential, nil
}

//todo modify the req type to []interface of credential , fetched from level db outside
func (ontVdri *OntVDRI) PresentProof(req *message.RequestPresentation) (*message.Presentation, error) {

	//fixme
	presentation := new(message.Presentation)
	presentation.Type = vdri.PresentationProofSpec
	presentation.Comment = "sample presentation"
	presentation.Connection = service.ReverseConnection(req.Connection)
	presentation.Thread = message.Thread{
		ID: req.Id,
	}

	return presentation, nil
}
func (o OntVDRI) GetDIDDoc(did string) (vdri.CommonDIDDoc, error) {

	//todo implement
	//only for test
	if did == "ontdid:ont:testdid1" {
		return service.DIDDoc{
			Context:        nil,
			Id:             "",
			PublicKey:      nil,
			Authentication: nil,
			Controller:     nil,
			Recovery:       nil,
			Service: []service.ServiceDoc{{
				ServiceID:       "ontdid:ont:serviceid1",
				ServiceType:     "ontdid",
				ServiceEndpoint: "http://192.168.1.114:8080",
			}},
			Attribute: nil,
			Created:   nil,
			Updated:   nil,
			Proof:     nil,
		}, nil
	}
	if did == "ontdid:ont:testdid2" {
		return service.DIDDoc{
			Context:        nil,
			Id:             "",
			PublicKey:      nil,
			Authentication: nil,
			Controller:     nil,
			Recovery:       nil,
			Service: []service.ServiceDoc{{
				ServiceID:       "ontdid:ont:serviceid2",
				ServiceType:     "ontdid",
				ServiceEndpoint: "http://0.0.0.0:8080",
				//ServiceEndpoint: "http://192.168.2.235:8080",
			}},
			Attribute: nil,
			Created:   nil,
			Updated:   nil,
			Proof:     nil,
		}, nil
	}
	return service.DIDDoc{
		Context:        nil,
		Id:             "",
		PublicKey:      nil,
		Authentication: nil,
		Controller:     nil,
		Recovery:       nil,
		Service: []service.ServiceDoc{{
			ServiceID:       "ontdid:ont:serviceid",
			ServiceType:     "ontdid",
			ServiceEndpoint: "http://0.0.0.0:8080",
		}},
		Attribute: nil,
		Created:   nil,
		Updated:   nil,
		Proof:     nil,
	}, nil
}
