package vdri

import "git.ont.io/ontid/otf/service"

type OntVDRI struct {
}

func (o OntVDRI) GetDIDDoc(did string) CommonDIDDoc {

	//todo
	return service.ServiceDoc{
		ServiceID:       "",
		ServiceType:     "",
		ServiceEndpoint: "",
	}
}
