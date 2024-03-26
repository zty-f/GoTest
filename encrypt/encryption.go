package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
)

// 私钥生成
var privateKey = []byte(`
-----BEGIN RSA PRIVATE KEY-----
MIGrAgEAAiEA4AAWS8HpsopKRXgoIgH59GrftgCBmtbTKVqcRe5b8KkCAwEAAQIh
AKgz6IoippYG+haIT7qZuVLDsYg2Y2NP9KkIUU2QU8jxAhEA4CXa61AGOlpVkVqT
9nwWWwIRAP/U3W87T02sJ8MLMH4pbEsCEAkJPaQS281qatyrPB/JrNUCEQCPPUn2
O4j9fkSNCjjOirbdAhB4QSuToYz78eTgSo1KqL++
-----END RSA PRIVATE KEY-----
`)

// 公钥: 根据私钥生成
var publicKey = []byte(`
-----BEGIN RSA PUBLIC KEY-----
MCgCIQDgABZLwemyikpFeCgiAfn0at+2AIGa1tMpWpxF7lvwqQIDAQAB
-----END RSA PUBLIC KEY-----
`)

// 生成密钥
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

// 公钥加密
func RsaEncryptPublic(origData []byte) ([]byte, error) {
	//解密pem格式的公钥
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errors.New("public key error")
	}
	// 解析公钥
	pubInterface, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	//加密
	return rsa.EncryptPKCS1v15(rand.Reader, pubInterface, origData)
}

// 解密
func RsaDecrypt(ciphertext []byte) ([]byte, error) {
	//解密
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("private key error!")
	}
	//解析PKCS1格式的私钥
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// 解密
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}
func main() {
	//private, public, err := GenerateKey(256)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//privateKey = private
	//fmt.Println(string(private))
	//publicKey = public
	//fmt.Println(string(public))
	data, _ := RsaEncryptPublic([]byte("2100051684"))
	fmt.Println(base64.StdEncoding.EncodeToString(data))
	origData, _ := RsaDecrypt(data)
	fmt.Println(string(origData))
}
