package did

type Did interface {
	ValidateDid(did string) bool
	NewDid() string
	GetDidType() string
	String() string
}
