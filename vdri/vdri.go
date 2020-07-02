package vdri

import "git.ont.io/ontid/otf/message"

type VDRI interface {
	GetDIDDoc(did string) (CommonDIDDoc, error)
	OfferCredential(req *message.ProposalCredential) (*message.OfferCredential, error)
	IssueCredential(req *message.RequestCredential) (*message.IssueCredential, error)
	PresentProof(req *message.RequestPresentation) (*message.Presentation, error)
}

type CommonDIDDoc interface {
	GetServicePoint(serviceid string) (string, error)
}
