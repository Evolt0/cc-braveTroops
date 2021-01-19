package RSA

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
)

// sign 签名 userPub 用户公钥 message 需要验证的信息
func Verify(sign string, userPub string, message string) error {
	PublicKey, err := parsingRsaPublicKey(userPub) // 解密公匙
	if err != nil {
		return err
	}
	decodeString, err := base64.StdEncoding.DecodeString(sign)
	if err != nil {
		return err
	}
	hashed := sha256.Sum256([]byte(message))
	err = rsa.VerifyPKCS1v15(PublicKey, crypto.SHA256, hashed[:], decodeString)
	if err != nil {
		return err
	}
	return nil
}

// 解析公匙
func parsingRsaPublicKey(userPub string) (*rsa.PublicKey, error) {
	// pem解码
	DEUserPub, _ := pem.Decode([]byte(userPub))
	if DEUserPub == nil {
		return nil, fmt.Errorf("failed to decode userPub")
	}
	// der解码，最终返回一个公匙对象
	pubKey, err := x509.ParsePKCS1PublicKey(DEUserPub.Bytes)
	if err != nil {
		return nil, err
	}
	return pubKey, nil
}
