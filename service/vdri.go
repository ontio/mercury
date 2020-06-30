package service

type VDRI interface {
	GetDIDDoc(did string) (CommonDIDDoc, error)
}

type CommonDIDDoc interface {
	GetServicePoint(serviceid string) (string, error)
}
