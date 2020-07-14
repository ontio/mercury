package ecdsa

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"git.ont.io/ontid/otf/packager"
	"git.ont.io/ontid/otf/utils"
	"github.com/ethereum/go-ethereum/crypto/ecies"
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

func (bp *Packager) PackMessage(envelope *packager.Envelope) ([]byte, error) {
	pub, err := utils.GetPubKeyByDid(envelope.ToDID, bp.ontSdk)
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
	data, err := Encrypt(pk, envelope.Message.Data)
	if err != nil {
		return nil, err
	}
	sign, err := bp.acct.Sign(data)
	if err != nil {
		return nil, err
	}
	packMsg := &packager.Envelope{
		Message: &packager.MessageData{
			Data: data,
			MsgType: envelope.Message.MsgType,
			Sign: sign,
		},
		FromDID: envelope.FromDID,
		ToDID:   envelope.ToDID,
	}
	jsonBytes, err := json.Marshal(packMsg)
	if err != nil {
		return nil, err
	}
	return jsonBytes, nil
}
func (bp *Packager) UnpackMessage(encMessage []byte) (*packager.Envelope, error) {
	data := &packager.Envelope{}
	err := json.Unmarshal(encMessage, data)
	if err != nil {
		return nil, err
	}
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
	sig, err := signature.Deserialize(data.Message.Sign)
	if err != nil {
		return nil, err
	}
	if !signature.Verify(pk, data.Message.Data, sig) {
		return nil, fmt.Errorf("data verify sign failed")
	}
	msg, err := Decrypt(bp.acct.PrivateKey, data.Message.Data)
	if err != nil {
		return nil, err
	}
	return &packager.Envelope{
		Message: &packager.MessageData{
			Data: msg,
			MsgType: data.Message.MsgType,
		},
		FromDID: data.FromDID,
		ToDID:   data.ToDID,
	}, nil
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
