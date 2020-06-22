package packager

import "crypto"

type Crypto interface {
	Encrypt(crypto.PublicKey, []byte) ([]byte, error)
	Decrypt(crypto.PrivateKey, []byte) ([]byte, error)
}
