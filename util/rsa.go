package util

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"test/constant"
)

// GenerateKey 生成rsa密钥对
func GenerateKey(bits int) ([]byte, []byte, error) {
	// 生成私钥文件
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, nil, err
	}
	// 生成公钥文件
	publicKey := &privateKey.PublicKey
	return pem.EncodeToMemory(
			&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(privateKey)}), pem.EncodeToMemory(
			&pem.Block{Type: "RSA PUBLIC KEY", Bytes: x509.MarshalPKCS1PublicKey(publicKey)}), nil
}

// RsaEncrypt 公钥加密
func RsaEncrypt(origData string) (string, error) {
	publicKey := []byte(constant.PublicKey)
	//解密pem格式的公钥
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return "", errors.New("public key error")
	}
	// 解析公钥
	pubInterface, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		return "", err
	}
	//加密
	bytes, err := rsa.EncryptPKCS1v15(rand.Reader, pubInterface, []byte(origData))
	if err != nil {
		return "", err
	}
	// base64编码
	return base64.StdEncoding.EncodeToString(bytes), nil
}

// RsaDecrypt 私钥解密
func RsaDecrypt(ciphertext string) (string, error) {
	// base64解码
	cipherBytes, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}
	privateKey := []byte(constant.PrivateKey)
	//解密
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return "", errors.New("private key error")
	}
	//解析PKCS1格式的私钥
	private, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}
	// 解密
	bytes, err := rsa.DecryptPKCS1v15(rand.Reader, private, cipherBytes)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
