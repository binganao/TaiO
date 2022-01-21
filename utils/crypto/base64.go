package crypto

import (
	"encoding/base64"
	"log"
)

func Base64Enctypto(raw string) string {
	input := []byte(raw)
	encodeString := base64.StdEncoding.EncodeToString(input)
	return encodeString
}

func Base64Decrypto(raw string) string {
	decodeBytes, err := base64.StdEncoding.DecodeString(raw)
	if err != nil {
		log.Fatalln(err)
	}
	return string(decodeBytes)
}
