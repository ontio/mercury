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
	"strconv"

	"github.com/ontio/mercury/common/message"
	sdk "github.com/ontio/ontology-go-sdk"
)

type PublicKey struct {
	ID           string
	Type         string
	Controller   string
	PublicKeyHex string
}

type Authentication struct {
	Did       string
	PublicKey PublicKey
}

type ServiceDoc struct {
	ServiceID       string
	ServiceType     string
	ServiceEndpoint string
}

type Doc struct {
	Context        []string
	Id             string
	PublicKey      []PublicKey
	Authentication Authentication
	Controller     interface{}
	Recovery       interface{}
	Service        []ServiceDoc
	Attribute      interface{}
	Created        interface{}
	Updated        interface{}
	Proof          interface{}
}

type DidPubkey struct {
	Id           string      `json:"id"`
	Type         string      `json:"type"`
	Controller   interface{} `json:"controller"`
	PublicKeyHex string      `json:"publicKeyHex"`
}

func ValidateDid(did string) bool {
	return sdk.VerifyID(CutDId(did))
}

func GetDidDocByDid(did string, ontSdk *sdk.OntologySdk) (*message.DIDDoc, error) {
	if ontSdk.Native == nil || ontSdk.Native.OntId == nil {
		return nil, fmt.Errorf("ontsdk is nil")
	}
	data, err := ontSdk.Native.OntId.GetDocumentJson(CutDId(did))
	if err != nil {
		return nil, err
	}
	doc := &message.DIDDoc{}
	err = json.Unmarshal([]byte(string(data)), doc)
	if err != nil {
		return nil, err
	}
	return doc, nil
}

func GetServiceEndpointByDid(did string, ontSdk *sdk.OntologySdk) ([]string, error) {
	if ontSdk.Native == nil || ontSdk.Native.OntId == nil {
		return nil, fmt.Errorf("ontsdk is nil")
	}
	data, err := ontSdk.Native.OntId.GetDocumentJson(CutDId(did))
	if err != nil {
		return nil, err
	}
	doc := &message.DIDDoc{}
	err = json.Unmarshal([]byte(string(data)), &doc)
	if err != nil {
		return nil, err
	}
	addrs := make([]string, 0)
	for _, endPoint := range doc.Service {
		addrs = append(addrs, endPoint.ServiceEndpoint)
	}
	return addrs, nil
}

func GetPubKeyByDid(did string, ontSdk *sdk.OntologySdk) (string, error) {
	if ontSdk.Native == nil || ontSdk.Native.OntId == nil {
		return "", fmt.Errorf("ontsdk is nil")
	}
	index, err := strconv.ParseInt(GetIndex(did), 10, 32)
	if err != nil {
		return "", err
	}
	pubKey, err := ontSdk.Native.OntId.GetPublicKeysJson(CutDId(did))
	if err != nil {
		return "", err
	}
	var pks []DidPubkey
	err = json.Unmarshal(pubKey, &pks)
	if err != nil {
		return "", err
	}
	if len(pks) < int(index) {
		return "", fmt.Errorf("no public key found")
	}
	return pks[index-1].PublicKeyHex, nil
}
