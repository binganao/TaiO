package crypto

import (
	"encoding/base64"
	"github.com/binganao/TaiO/pkg/logger"
)

func Base64Enctypto(raw string) string {
	input := []byte(raw)
	encodeString := base64.StdEncoding.EncodeToString(input)
	return encodeString
}

func Base64Decrypto(raw string) string {
	decodeBytes, err := base64.StdEncoding.DecodeString(raw)
	if err != nil {
		logger.Info("输入信息解码失败")
		return ""
	}
	return string(decodeBytes)
}
