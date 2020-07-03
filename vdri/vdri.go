package vdri

import (
	"git.ont.io/ontid/otf/message"
	"git.ont.io/ontid/otf/store"
)

type VDRI interface {
	GetDIDDoc(did string) (CommonDIDDoc, error)
	OfferCredential(req *message.ProposalCredential) (*message.OfferCredential, error)
	IssueCredential(req *message.RequestCredential) (*message.IssueCredential, error)
	PresentProof(req *message.RequestPresentation, db store.Store) (*message.Presentation, error)
}

type CommonDIDDoc interface {
	GetServicePoint(serviceid string) (string, error)
}
