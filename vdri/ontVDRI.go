package vdri

import "git.ont.io/ontid/otf/service"

//todo use sdk query from smart contract
type OntVDRI struct {
}

func (o OntVDRI) GetDIDDoc(did string) CommonDIDDoc {

	//todo implement
	//only for test
	if did == "did:ont:testdid1" {
		return service.DIDDoc{
			Context:        nil,
			Id:             "",
			PublicKey:      nil,
			Authentication: nil,
			Controller:     nil,
			Recovery:       nil,
			Service:        []service.ServiceDoc{{
				ServiceID:"did:ont:serviceid1",
				ServiceType:"ontdid",
				ServiceEndpoint:"http://0.0.0.0:8080",
			}},
			Attribute:      nil,
			Created:        nil,
			Updated:        nil,
			Proof:          nil,
		}
	}
	if did == "did:ont:testdid2" {
		return service.DIDDoc{
			Context:        nil,
			Id:             "",
			PublicKey:      nil,
			Authentication: nil,
			Controller:     nil,
			Recovery:       nil,
			Service:         []service.ServiceDoc{{
				ServiceID:"did:ont:serviceid2",
				ServiceType:"ontdid",
				ServiceEndpoint:"http://0.0.0.0:8089",
			}},
			Attribute:      nil,
			Created:        nil,
			Updated:        nil,
			Proof:          nil,
		}
	}
}
