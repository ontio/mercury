package vdri

type DidDoc interface {
	GetServicePoint(serviceId string) (string, error)
	GetServiceEndpointByDid(did string, sdk interface{}) ([]string, error)
	GetDidDocByDid(did string, sdk interface{}) (interface{}, error)
}
