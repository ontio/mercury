package utils

import "encoding/base64"

func Base64Encode(bs []byte)string{
	return base64.StdEncoding.EncodeToString(bs)
}

func Base64Decode(s string)([]byte,error){
	return base64.StdEncoding.DecodeString(s)
}