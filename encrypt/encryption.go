package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"time"
)

// 私钥生成
var privateKey = []byte(`
-----BEGIN RSA PRIVATE KEY-----
MIGCAgEAAhg9XZlJScmq35D60n+C+cNZUUsbVE0jq0ECAwEAAQIYMugVxH2I5c6n
N9wtCabsiejwkFG29AC9Agx9Bh4VzHwG1GxdCFcCDH2nG9vbvUoLxfwKJwIMYIXp
kkDV/Fvh8Y1vAgwI/29yRgD/DWrHCp8CDG0r2zc4jgl4BN0+dg==
-----END RSA PRIVATE KEY-----
`)

// 公钥: 根据私钥生成
var publicKey = []byte(`
-----BEGIN RSA PUBLIC KEY-----
MB8CGD1dmUlJyarfkPrSf4L5w1lRSxtUTSOrQQIDAQAB
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
	t := time.Now()
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
	fmt.Println("耗时：", time.Since(t).Microseconds())
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}
func main() {
	//private, public, err := GenerateKey(190)
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
	str := base64.StdEncoding.EncodeToString(data)
	decodeString, _ := base64.StdEncoding.DecodeString(str)
	origData, err := RsaDecrypt(decodeString)
	fmt.Println("-----------------")
	fmt.Println(err)
	fmt.Println(string(origData))
	fmt.Println("-----------------")
	//origData, err = RsaDecrypt([]byte(str))
	//fmt.Println(err)
	//fmt.Println(string(origData))
	//fmt.Print("-------------------")
}
