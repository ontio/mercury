package crypto

type Crypto interface {
	Encrypt([]byte,[]byte) string
	Decrypt([]byte,[]byte) string
}