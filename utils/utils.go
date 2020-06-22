package utils

import (
	"fmt"
	sdk "github.com/ontio/ontology-go-sdk"
	"github.com/howeyc/gopass"
)

var Version = ""


func OpenAccount(path string) (*sdk.Account,error){
	ontSdk := sdk.NewOntologySdk()
	wallet,err := ontSdk.OpenWallet("./wallet.dat")
	if err != nil {
		return nil,err
	}
	pwd, err := GetPassword()
	if err != nil {
		return nil,err
	}
	defer ClearPasswd(pwd)
	account, err :=wallet.GetDefaultAccount(pwd)
	if err != nil {
		return nil,err
	}
	return account,nil
}

func GetPassword() ([]byte, error) {
	fmt.Printf("Password:")
	passwd, err := gopass.GetPasswd()
	if err != nil {
		return nil, err
	}
	return passwd, nil
}

func ClearPasswd(passwd []byte) {
	size := len(passwd)
	for i := 0; i < size; i++ {
		passwd[i] = 0
	}
}