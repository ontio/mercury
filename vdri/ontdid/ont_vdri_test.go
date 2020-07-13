package ontdid

import (
	"fmt"
	ontology_go_sdk "github.com/ontio/ontology-go-sdk"
	"github.com/ontio/ontology/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

var testOntSdk *ontology_go_sdk.OntologySdk
var testDefAcc *ontology_go_sdk.Account

func Init() {
	testOntSdk = ontology_go_sdk.NewOntologySdk()
	testOntSdk.NewRpcClient().SetAddress("http://polaris2.ont.io:20336")

	var err error
	var wallet *ontology_go_sdk.Wallet
	if !common.FileExisted("./wallet.dat") {
		wallet, err = testOntSdk.CreateWallet("./wallet.dat")
		if err != nil {
			fmt.Println("[CreateWallet] error:", err)
			return
		}
	} else {
		wallet, err = testOntSdk.OpenWallet("./wallet.dat")
		if err != nil {
			fmt.Println("[CreateWallet] error:", err)
			return
		}
	}
	_, err = wallet.NewDefaultSettingAccount([]byte("123456"))
	if err != nil {
		fmt.Println("")
		return
	}
	wallet.Save()
	testWallet, err := testOntSdk.OpenWallet("./wallet.dat")
	if err != nil {
		fmt.Printf("account.Open error:%s\n", err)
		return
	}
	testDefAcc, err = testWallet.GetDefaultAccount([]byte("123456"))
	if err != nil {
		fmt.Printf("GetDefaultAccount error:%s\n", err)
		return
	}

	return

}

func TestOntVDRI_VerifyCred(t *testing.T) {
	Init()
	s := "eyJhbGciOiJFUzI1NiIsImtpZCI6ImRpZDpvbnQ6VEtnSDZKaVlXU0x4V3BDeW9EWnVreTZycE5yRzc5emVkeiNrZXlzLTEiLCJ0eXAiOiJKV1QifQ==.eyJpc3MiOiJkaWQ6b250OlRLZ0g2SmlZV1NMeFdwQ3lvRFp1a3k2cnBOckc3OXplZHoiLCJleHAiOjE1OTQ3MDc2MTksIm5iZiI6MTU5NDYyMTIyMCwiaWF0IjoxNTk0NjIxMjIwLCJqdGkiOiJ1cm46dXVpZDpiYzg5MTM4Ni1jNWRhLTRjZGUtODdiMi05NTdhYjVmMmZjNGEiLCJ2YyI6eyJAY29udGV4dCI6WyJodHRwczovL3d3dy53My5vcmcvMjAxOC9jcmVkZW50aWFscy92MSIsImh0dHBzOi8vb250aWQub250LmlvL2NyZWRlbnRpYWxzL3YxIiwiY29udGV4dDEiLCJjb250ZXh0MiJdLCJ0eXBlIjpbIlZlcmlmaWFibGVDcmVkZW50aWFsIiwib3RmIl0sImNyZWRlbnRpYWxTdWJqZWN0IjpbeyJuYW1lIjoiYWdlIiwidmFsdWUiOiJncmVhdGVyIHRoYW4gMTgifV0sImNyZWRlbnRpYWxTdGF0dXMiOnsiaWQiOiIwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwIiwidHlwZSI6IkF0dGVzdENvbnRyYWN0In0sInByb29mIjp7ImNyZWF0ZWQiOiIyMDIwLTA3LTEzVDA2OjIwOjIwWiIsInByb29mUHVycG9zZSI6ImFzc2VydGlvbk1ldGhvZCJ9fX0=.lKbIq3yLgLN8oDljyhIdmzCmtuIppWilN/iycK4JNciApKSwH98K4EIa6fNGQaGS+M8+nOmqNcwM36aMGXPxnA=="

	err := testOntSdk.Credential.VerifyJWTIssuerSignature(s)
	assert.Nil(t, err, "VerifyJWTStatus failed")
	err = testOntSdk.Credential.VerifyJWTDate(s)
	assert.Nil(t, err, "VerifyJWTDate failed")

	jwtc, err := ontology_go_sdk.DeserializeJWT(s)
	assert.Nil(t, err, "Base64Decode failed")
	fmt.Printf("%v\n", jwtc.Payload.VC.CredentialSubject)

	tmp := jwtc.Payload.VC.CredentialSubject.([]interface{})
	m := tmp[0].(map[string]interface{})
	assert.Equal(t, m["value"], "greater than 18", "value not equal")
	assert.Equal(t, m["name"], "age", "age not equal")
}

func Test_ONTID(t *testing.T) {
	Init()
	bs, err := testOntSdk.Native.OntId.GetDocumentJson("did:ont:TKgH6JiYWSLxWpCyoDZuky6rpNrG79zedz")
	assert.Nil(t, err, "GetDocumentJson failed")
	fmt.Printf("%s\n", bs)
}

func Test_Presentastion(t *testing.T) {
	Init()
	s := "eyJhbGciOiJFUzI1NiIsImtpZCI6ImRpZDpvbnQ6VEtnSDZKaVlXU0x4V3BDeW9EWnVreTZycE5yRzc5emVkeiNrZXlzLTEiLCJ0eXAiOiJKV1QifQ==.eyJpc3MiOiJkaWQ6b250OlRLZ0g2SmlZV1NMeFdwQ3lvRFp1a3k2cnBOckc3OXplZHoiLCJhdWQiOiIiLCJqdGkiOiJ1cm46dXVpZDpjYzdiZGMwNC1iMGViLTQ1M2EtOTNmYy04YWY0NzNlNTI0NGYiLCJ2cCI6eyJAY29udGV4dCI6WyJodHRwczovL3d3dy53My5vcmcvMjAxOC9jcmVkZW50aWFscy92MSIsImh0dHBzOi8vb250aWQub250LmlvL2NyZWRlbnRpYWxzL3YxIiwiY29udGV4dDEiLCJjb250ZXh0MiJdLCJ0eXBlIjpbIlZlcmlmaWFibGVDcmVkZW50aWFsIiwib3RmIl0sInZlcmlmaWFibGVDcmVkZW50aWFsIjpbImV5SmhiR2NpT2lKRlV6STFOaUlzSW10cFpDSTZJbVJwWkRwdmJuUTZWRXRuU0RaS2FWbFhVMHg0VjNCRGVXOUVXblZyZVRaeWNFNXlSemM1ZW1Wa2VpTnJaWGx6TFRFaUxDSjBlWEFpT2lKS1YxUWlmUT09LmV5SnBjM01pT2lKa2FXUTZiMjUwT2xSTFowZzJTbWxaVjFOTWVGZHdRM2x2UkZwMWEzazJjbkJPY2tjM09YcGxaSG9pTENKbGVIQWlPakUxT1RRM01EYzJNVGtzSW01aVppSTZNVFU1TkRZeU1USXlNQ3dpYVdGMElqb3hOVGswTmpJeE1qSXdMQ0pxZEdraU9pSjFjbTQ2ZFhWcFpEcGlZemc1TVRNNE5pMWpOV1JoTFRSalpHVXRPRGRpTWkwNU5UZGhZalZtTW1aak5HRWlMQ0oyWXlJNmV5SkFZMjl1ZEdWNGRDSTZXeUpvZEhSd2N6b3ZMM2QzZHk1M015NXZjbWN2TWpBeE9DOWpjbVZrWlc1MGFXRnNjeTkyTVNJc0ltaDBkSEJ6T2k4dmIyNTBhV1F1YjI1MExtbHZMMk55WldSbGJuUnBZV3h6TDNZeElpd2lZMjl1ZEdWNGRERWlMQ0pqYjI1MFpYaDBNaUpkTENKMGVYQmxJanBiSWxabGNtbG1hV0ZpYkdWRGNtVmtaVzUwYVdGc0lpd2liM1JtSWwwc0ltTnlaV1JsYm5ScFlXeFRkV0pxWldOMElqcGJleUp1WVcxbElqb2lZV2RsSWl3aWRtRnNkV1VpT2lKbmNtVmhkR1Z5SUhSb1lXNGdNVGdpZlYwc0ltTnlaV1JsYm5ScFlXeFRkR0YwZFhNaU9uc2lhV1FpT2lJd01EQXdNREF3TURBd01EQXdNREF3TURBd01EQXdNREF3TURBd01EQXdNREF3TURBd01EQXdJaXdpZEhsd1pTSTZJa0YwZEdWemRFTnZiblJ5WVdOMEluMHNJbkJ5YjI5bUlqcDdJbU55WldGMFpXUWlPaUl5TURJd0xUQTNMVEV6VkRBMk9qSXdPakl3V2lJc0luQnliMjltVUhWeWNHOXpaU0k2SW1GemMyVnlkR2x2YmsxbGRHaHZaQ0o5ZlgwPS5sS2JJcTN5TGdMTjhvRGxqeWhJZG16Q210dUlwcFdpbE4vaXljSzRKTmNpQXBLU3dIOThLNEVJYTZmTkdRYUdTK004K25PbXFOY3dNMzZhTUdYUHhuQT09Il0sInByb29mIjp7ImNyZWF0ZWQiOiIyMDIwLTA3LTEzVDA3OjEzOjM3WiIsInByb29mUHVycG9zZSI6ImFzc2VydGlvbk1ldGhvZCJ9fX0=.tejyKEZ88UrLqt/elK/tQ6DdPOQW+kBjfYh6T3h/AezeLLi4dMJO0Mx8mt4V0X5RBMCvGGaPyLFiZCwCIRTDRQ=="
	err := testOntSdk.Credential.VerifyJWTIssuerSignature(s)
	assert.Nil(t, err, "VerifyJWTIssuerSignature failed")

	err = testOntSdk.Credential.VerifyJWTDate(s)
	assert.Nil(t, err, "VerifyJWTDate failed")

	presentation, err := testOntSdk.Credential.JWTPresentation2Json(s)
	assert.Nil(t, err, "JWTPresentation2Json failed")

	vc := presentation.VerifiableCredential
	assert.Equal(t, len(vc), 1, "vc len is not 1")
	tmp := vc[0].CredentialSubject.([]interface{})
	m := tmp[0].(map[string]interface{})
	assert.Equal(t, m["value"], "greater than 18", "value not equal")
	assert.Equal(t, m["name"], "age", "age not equal")

}
