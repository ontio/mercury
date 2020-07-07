package ontdid

import (
	"encoding/json"
	"fmt"
	"git.ont.io/ontid/otf/message"
	"git.ont.io/ontid/otf/service"
	"git.ont.io/ontid/otf/service/controller"
	"git.ont.io/ontid/otf/store"
	"git.ont.io/ontid/otf/utils"
	"git.ont.io/ontid/otf/vdri"
	"github.com/google/uuid"
	sdk "github.com/ontio/ontology-go-sdk"
	"time"
)

var (
	contexts = []string{"context1", "context2"}
	types    = []string{"otf"}
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

	expirationDate := time.Now().UTC().Unix() + 86400

	vc, err := ontVdri.ontSdk.Credential.CreateJWTCredential(contexts, types, subs, req.Connection.TheirDid, expirationDate, "", nil, ontVdri.acct)
	if err != nil {
		fmt.Printf("CreateJWTCredential err:%s\n", err.Error())
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

func (ontVdri *OntVDRI) PresentProof(req *message.RequestPresentation, db store.Store) (*message.Presentation, error) {

	holderdid := req.Connection.TheirDid
	creds := make([]string, 0)
	for _, attachment := range req.RequestPresentationAttach {
		b64 := attachment.Data.Base64
		//should be cred id
		bts, err := utils.Base64Decode(b64)
		if err != nil {
			return nil, err
		}
		credid := string(bts)

		key := []byte(fmt.Sprintf("%s_%s", controller.CredentialKey, credid))
		data, err := db.Get(key)
		if err != nil {
			return nil, err
		}

		cred := new(message.IssueCredential)
		err = json.Unmarshal(data, cred)
		if err != nil {
			return nil, err
		}

		s := cred.CredentialsAttach[0].Data.Base64
		creds = append(creds, s)
	}

	presentation := new(message.Presentation)
	ps, err := ontVdri.ontSdk.Credential.CreateJWTPresentation(creds, contexts, types, holderdid, "", "", ontVdri.acct)
	if err != nil {
		return nil, err
	}
	presentation.Type = vdri.PresentationProofSpec
	presentation.Id = utils.GenUUID()
	presentation.Connection = service.ReverseConnection(req.Connection)
	presentation.Formats = []message.Format{message.Format{
		AttachID: "1", //magic index
		Format:   "base64",
	}}
	presentation.PresentationAttach = []message.Attachment{
		{
			Id:          "1",
			LastModTime: time.Now(),
			Data: message.Data{
				Base64: ps,
			},
		},
	}
	presentation.Thread = message.Thread{
		ID: req.Id,
	}

	return presentation, nil
}
func (o OntVDRI) GetDIDDoc(did string) (vdri.CommonDIDDoc, error) {

	bts, err := o.ontSdk.Native.OntId.GetDocumentJson(did)
	if err != nil {
		return nil, err
	}

	doc := new(message.DIDDoc)
	err = json.Unmarshal(bts, doc)
	if err != nil {
		return nil, err
	}
	return doc, nil

	/*	//todo implement
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
		}, nil*/
}
