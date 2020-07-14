package utils

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/howeyc/gopass"
	sdk "github.com/ontio/ontology-go-sdk"
	"strings"
)

var Version = ""

func OpenAccount(path string, ontSdk *sdk.OntologySdk) (*sdk.Account, error) {
	wallet, err := ontSdk.OpenWallet(path)
	if err != nil {
		return nil, err
	}
	pwd, err := GetPassword()
	if err != nil {
		return nil, err
	}
	defer ClearPasswd(pwd)
	account, err := wallet.GetDefaultAccount(pwd)
	if err != nil {
		return nil, err
	}
	return account, nil
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

func GenUUID() string {
	return uuid.New().String()
}

func CutDId(did string)string {
	var realdid string
	if strings.Contains(did, "#") {
		realdid = strings.Split(did, "#")[0]
	} else {
		realdid = did
	}
	return realdid

}