package BTCUtils

import (
	"crypto/sha256"
	"encoding/base64"
)

func Sha256Hassh(input []byte)[]byte{
	sha256hsah := sha256.New()
	sha256hsah.Write(input)
	hashBytes := sha256hsah.Sum(nil)
	return hashBytes
}

//Base64编码工具
func Base64Encode(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}
func Base64Decode(str string) string {
	result, _ := base64.StdEncoding.DecodeString(str)
	return string(result)
}

