package ontdid

import (
	"git.ont.io/ontid/otf/message"
	"git.ont.io/ontid/otf/service"
	"git.ont.io/ontid/otf/vdri"
)

//todo use sdk query from smart contract
type OntVDRI struct {
}

func NewOntVDRI() *OntVDRI {
	return &OntVDRI{}
}

func (ontVdri *OntVDRI) OfferCredential(req message.ProposalCredential) (*message.OfferCredential, error) {
	return nil, nil
}
func (ontVdri *OntVDRI) IssueCredential(req message.RequestCredential) (*message.IssueCredential, error) {
	return nil, nil
}

func (ontVdri *OntVDRI) PresentProof(req message.RequestPresentation) (*message.Presentation, error) {
	return nil, nil
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
