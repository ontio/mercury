package did

import (
	"github.com/google/uuid"
	sdk "github.com/ontio/ontology-go-sdk"
)

type OntDID struct{
	account *sdk.Account
}

func (o *OntDID)ValidateDid(did string)bool{
	return true
}

func (o *OntDID)NewDid()string{

	//fixme
	return "did:ont:testdid" + uuid.New().String()
}

func(o *OntDID)GetDidType()string{
	return "test"
}

func(o *OntDID)String() string{
	//fixme
	return "did:ont:" + o.account.Address.ToBase58()
}