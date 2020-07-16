package utils

import (
	"encoding/json"
	"fmt"
	"git.ont.io/ontid/otf/message"
	"git.ont.io/ontid/otf/store"
	"github.com/google/uuid"
	"github.com/howeyc/gopass"
	sdk "github.com/ontio/ontology-go-sdk"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

var Version = ""

const (
	InvitationKey    = "Invitation"
	ConnectionReqKey = "ConnectionReq"
	ConnectionKey    = "Connection"
	GeneralMsgKey    = "General"

	ACK_SUCCEED = "succeed"
	ACK_FAILED  = "failed"
)

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

func CutDId(did string) string {
	var realdid string
	if strings.Contains(did, "#") {
		realdid = strings.Split(did, "#")[0]
	} else {
		realdid = did
	}
	return realdid

}

func CheckConnection(mydid, theirdid string, db store.Store) error {
	connectionKey := []byte(fmt.Sprintf("%s_%s", ConnectionKey, mydid))
	data, err := db.Get(connectionKey)
	if err != nil {
		return nil
	}

	cr := new(message.ConnectionRec)
	err = json.Unmarshal(data, cr)
	if err != nil {
		return err
	}
	_, ok := cr.Connections[theirdid]
	if !ok {
		return fmt.Errorf("connection not found!")
	}
	return nil
}

func HttpPostData(client *http.Client, url, data string) ([]byte, error) {
	resp, err := client.Post(url, "application/json", strings.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("http post request:%s error:%s", data, err)
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func NewClient() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost:   5,
			DisableKeepAlives:     false,
			IdleConnTimeout:       time.Second * 300,
			ResponseHeaderTimeout: time.Second * 300,
		},
		Timeout: time.Second * 300,
	}
}
