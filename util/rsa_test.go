package util

import (
	"fmt"
	"github.com/spf13/cast"
	"testing"
)

func TestGenerateKey(t *testing.T) {
	privateKey, publicKey, err := GenerateKey(190)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(privateKey)
	fmt.Println(publicKey)
}

func TestRsa(t *testing.T) {
	stuId := 2100051684
	encrypt, err := RsaEncrypt(cast.ToString(stuId))
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(encrypt)
	decrypt, err := RsaDecrypt(encrypt)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(decrypt)
	fmt.Println(cast.ToInt(decrypt) == stuId)
}
