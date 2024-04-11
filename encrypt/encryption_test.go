package main

import (
	"github.com/spf13/cast"
	"github.com/stretchr/testify/assert"
	"testing"
)

func BenchmarkRsaDecrypt(b *testing.B) {
	stuId := 2100051671
	encrypt, err := RsaEncryptPublic([]byte(cast.ToString(stuId)))
	assert.Nil(b, err)

	for i := 0; i < b.N; i++ {
		decrypt, errInner := RsaDecrypt([]byte(encrypt))
		assert.Nil(b, errInner)
		assert.NotEqual(b, "", decrypt)
	}
}
