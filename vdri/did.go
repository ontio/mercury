package vdri


type Did interface {
	ValidateDid(did string) bool
	NewDid() string
	GetDidType() string
	String() string
}