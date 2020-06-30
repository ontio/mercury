package vdri

type VDRI interface {
	GetDIDDoc(did string) CommonDIDDoc
}

type CommonDIDDoc interface {
	GetServicePoint() string
}
