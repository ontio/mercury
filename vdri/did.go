package vdri

type Did interface {
	NewDid() string
	ValidateDid(did string) bool
}
