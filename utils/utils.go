/*
 * Copyright (C) 2018 The ontology Authors
 * This file is part of The ontology library.
 *
 * The ontology is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The ontology is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Lesser General Public License for more details.
 *
 * You should have received a copy of the GNU Lesser General Public License
 * along with The ontology.  If not, see <http://www.gnu.org/licenses/>.
 */

package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/howeyc/gopass"
	"github.com/ontio/mercury/common/message"
	"github.com/ontio/mercury/store"
	sdk "github.com/ontio/ontology-go-sdk"
)

var Version = ""

const (
	InvitationKey    = "Invitation"
	ConnectionReqKey = "ConnectionReq"
	ConnectionKey    = "Connection"
	BasicMsgKey      = "Basic"
	ACK_SUCCEED      = "succeed"
	ACK_FAILED       = "failed"
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
	index := strings.LastIndex(did, "@")
	if index != -1 {
		return did[:index]
	}
	return did
}

//Did format  did@index#svrIndex
func GetIndex(did string) string {
	index := strings.LastIndex(did, "@")
	svrIndex := strings.LastIndex(did, "#")
	if index != -1 {
		if svrIndex != -1 {
			return did[index+1 : svrIndex]
		} else {
			return did[index+1:]
		}
	}
	return "1" //default 1
}

func CheckConnection(myDid, theirDid string, db store.Store) error {
	connectionKey := []byte(fmt.Sprintf("%s_%s", ConnectionKey, myDid))
	data, err := db.Get(connectionKey)
	if err != nil {
		return err
	}

	cr := new(message.ConnectionRec)
	err = json.Unmarshal(data, cr)
	if err != nil {
		return err
	}
	_, ok := cr.Connections[theirDid]
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
