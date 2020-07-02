package service

//todo use sdk query from smart contract
type OntVDRI struct {
}

func NewOntVDRI() *OntVDRI {
	return &OntVDRI{}
}

func (o OntVDRI) GetDIDDoc(did string) (CommonDIDDoc, error) {

	//todo implement
	//only for test
	if did == "ontdid:ont:testdid1" {
		return DIDDoc{
			Context:        nil,
			Id:             "",
			PublicKey:      nil,
			Authentication: nil,
			Controller:     nil,
			Recovery:       nil,
			Service: []ServiceDoc{{
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
		return DIDDoc{
			Context:        nil,
			Id:             "",
			PublicKey:      nil,
			Authentication: nil,
			Controller:     nil,
			Recovery:       nil,
			Service: []ServiceDoc{{
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
	return DIDDoc{
		Context:        nil,
		Id:             "",
		PublicKey:      nil,
		Authentication: nil,
		Controller:     nil,
		Recovery:       nil,
		Service: []ServiceDoc{{
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
