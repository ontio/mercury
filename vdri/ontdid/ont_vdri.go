package ontdid

import (
	"git.ont.io/ontid/otf/message"
	"git.ont.io/ontid/otf/service"
	"git.ont.io/ontid/otf/vdri"
	"github.com/google/uuid"
)

//todo use sdk query from smart contract
type OntVDRI struct {
}

func NewOntVDRI() *OntVDRI {
	return &OntVDRI{}
}

func (ontVdri *OntVDRI) OfferCredential(req *message.ProposalCredential) (*message.OfferCredential, error) {
	return nil, nil
}
func (ontVdri *OntVDRI) IssueCredential(req *message.RequestCredential) (*message.IssueCredential, error) {
	//fixme
	credential := &message.IssueCredential{
		Type:              vdri.IssueCredentialSpec,
		Id:                uuid.New().String(),
		Comment:           "",
		Formats:           nil,
		CredentialsAttach: nil,
		Connection:        service.ReverseConnection(req.Connection),
		Thread: message.Thread{
			ID: req.Id,
		},
	}

	return credential, nil
}

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
