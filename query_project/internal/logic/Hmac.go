package logic

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
)

func ToHmac(salt, data string) string {
	hasher := hmac.New(sha256.New, []byte(salt))
	hasher.Write([]byte(data))
	hashedPassWord := hasher.Sum(nil)
	// 特殊字符mysql没法作为string存储，所以转成base64
	return base64.StdEncoding.EncodeToString(hashedPassWord)
}
