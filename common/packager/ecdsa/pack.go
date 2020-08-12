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

package ecdsa

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/crypto/ecies"
	"github.com/ontio/mercury/common/packager"
	"github.com/ontio/mercury/utils"
	"github.com/ontio/ontology-crypto/ec"
	"github.com/ontio/ontology-crypto/keypair"
	"github.com/ontio/ontology-crypto/signature"
	"github.com/ontio/ontology-crypto/sm2"
	sdk "github.com/ontio/ontology-go-sdk"
)

type Packager struct {
	ontSdk *sdk.OntologySdk
	acct   *sdk.Account
}

func New(ontSdk *sdk.OntologySdk, acct *sdk.Account) *Packager {
	return &Packager{
		ontSdk: ontSdk,
		acct:   acct,
	}
}

func (bp *Packager) PackConnection(connectionData []byte, toDid string) (*packager.MsgConnection, error) {
	pub, err := utils.GetPubKeyByDid(toDid, bp.ontSdk)
	if err != nil {
		return nil, err
	}
	pubKey, err := hex.DecodeString(pub)
	if err != nil {
		return nil, err
	}
	pk, err := keypair.DeserializePublicKey(pubKey)
	if err != nil {
		return nil, err
	}
	data, err := Encrypt(pk, connectionData)
	if err != nil {
		return nil, err
	}
	sign, err := bp.acct.Sign(data)
	if err != nil {
		return nil, err
	}
	return &packager.MsgConnection{
		Data: data,
		Sign: sign,
	}, nil
}

func (bp *Packager) UnPackConnection(data *packager.Envelope) (*packager.MsgConnection, error) {
	pub, err := utils.GetPubKeyByDid(data.FromDID, bp.ontSdk)
	if err != nil {
		return nil, err
	}
	pubKey, err := hex.DecodeString(pub)
	if err != nil {
		return nil, err
	}
	pk, err := keypair.DeserializePublicKey(pubKey)
	if err != nil {
		return nil, err
	}
	sig, err := signature.Deserialize(data.Connection.Sign)
	if err != nil {
		return nil, err
	}
	if !signature.Verify(pk, data.Connection.Data, sig) {
		return nil, fmt.Errorf("connection data verify sign failed")
	}
	msg, err := Decrypt(bp.acct.PrivateKey, data.Connection.Data)
	if err != nil {
		return nil, err
	}
	return &packager.MsgConnection{
		Data: msg,
	}, nil
}

func (bp *Packager) PackMessage(envelope *packager.MessageData, destDid string) (*packager.MessageData, error) {
	pub, err := utils.GetPubKeyByDid(destDid, bp.ontSdk)
	if err != nil {
		return nil, err
	}
	pubKey, err := hex.DecodeString(pub)
	if err != nil {
		return nil, err
	}
	pk, err := keypair.DeserializePublicKey(pubKey)
	if err != nil {
		return nil, err
	}
	data, err := Encrypt(pk, envelope.Data)
	if err != nil {
		return nil, err
	}
	sign, err := bp.acct.Sign(data)
	if err != nil {
		return nil, err
	}
	return &packager.MessageData{
		Data:    data,
		MsgType: envelope.MsgType,
		Sign:    sign,
	}, nil
}

func (bp *Packager) UnpackMessage(data *packager.MessageData, sourceDid string) (*packager.MessageData, error) {
	pub, err := utils.GetPubKeyByDid(sourceDid, bp.ontSdk)
	if err != nil {
		return nil, err
	}
	pubKey, err := hex.DecodeString(pub)
	if err != nil {
		return nil, err
	}
	pk, err := keypair.DeserializePublicKey(pubKey)
	if err != nil {
		return nil, err
	}
	sig, err := signature.Deserialize(data.Sign)
	if err != nil {
		return nil, err
	}
	if !signature.Verify(pk, data.Data, sig) {
		return nil, fmt.Errorf("data verify sign failed")
	}
	msg, err := Decrypt(bp.acct.PrivateKey, data.Data)
	if err != nil {
		return nil, err
	}
	return &packager.MessageData{
		Data: msg,
	}, nil
}

func (bp *Packager) PackData(envelope *packager.Envelope) ([]byte, error) {
	jsonBytes, err := json.Marshal(envelope)
	if err != nil {
		return nil, err
	}
	return jsonBytes, nil
}

func (bp *Packager) UnPackData(enData []byte) (*packager.Envelope, error) {
	data := &packager.Envelope{}
	err := json.Unmarshal(enData, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func Encrypt(pub keypair.PublicKey, m []byte) ([]byte, error) {
	switch key := pub.(type) {
	case *ec.PublicKey:
		if key.Algorithm == ec.SM2 {
			return sm2.Encrypt(key.PublicKey, m)
		} else if key.Algorithm == ec.ECDSA {
			pk := &ecies.PublicKey{
				X:      key.X,
				Y:      key.Y,
				Curve:  key.Curve,
				Params: ecies.ParamsFromCurve(key.Curve),
			}
			return ecies.Encrypt(rand.Reader, pk, m, nil, keypair.SerializePublicKey(key))
		} else {
			panic("unknown public key type")
		}
	default:
		return nil, errors.New("unsupported encryption key")
	}

}

func Decrypt(pri keypair.PrivateKey, c []byte) ([]byte, error) {
	switch key := pri.(type) {
	case *ec.PrivateKey:
		if key.Algorithm == ec.SM2 {
			return sm2.Decrypt(key.PrivateKey, c)
		} else if key.Algorithm == ec.ECDSA {
			sk := &ecies.PrivateKey{
				PublicKey: ecies.PublicKey{
					X:      key.X,
					Y:      key.Y,
					Curve:  key.Curve,
					Params: ecies.ParamsFromCurve(key.Curve),
				},
				D: key.D,
			}
			return sk.Decrypt(c, nil, keypair.SerializePublicKey(key.Public()))
		} else {
			panic("unknown private key type")
		}
	default:
		return nil, errors.New("unsupported decryption key")
	}
}
