package ecdsa

import (
	"bytes"
	"encoding/hex"
	"github.com/ontio/ontology-crypto/keypair"
	"testing"
)

func TestEncrypt(t *testing.T) {
	pub := "0375327258042b7c280bd5bf56d29bab6eca71ed8686da02b5615727cf3ad81c7a"
	pubKey, err := hex.DecodeString(pub)
	if err != nil {
		t.Fatalf("DecodeString err:%s", err)
	}
	pk, err := keypair.DeserializePublicKey(pubKey)
	if err != nil {
		t.Fatalf("deserialize failed:%s", err)
	}
	c, err := Encrypt(pk, []byte("data did"))
	if err != nil {
		t.Fatal(err)
	}
	pri := "120233cdefbae991896339e8982b71b20a79e207e9d4eb09add426d8f7f279dccae10375327258042b7c280bd5bf56d29bab6eca71ed8686da02b5615727cf3ad81c7a"
	priKey, err := hex.DecodeString(pri)
	if err != nil {
		t.Fatalf("DecodeString err:%s", err)
	}
	sk, err := keypair.DeserializePrivateKey(priKey)
	if err != nil {
		t.Fatalf("DeserializePrivateKey err:%s", err)
	}
	m, err := Decrypt(sk, c)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(m, []byte("data did")) {
		t.Fatal("decrypted message is wrong")
	}
}
