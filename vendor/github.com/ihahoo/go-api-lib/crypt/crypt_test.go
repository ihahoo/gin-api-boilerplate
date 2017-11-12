package crypt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMD5(t *testing.T) {
	assert.Equal(t, MD5("123456"), "e10adc3949ba59abbe56e057f20f883e", "MD5不匹配")
}

func TestSha1(t *testing.T) {
	assert.Equal(t, Sha1("123456"), "7c4a8d09ca3762af61e59520943dc26494f8941b", "sha1不匹配")
}

func TestSha256(t *testing.T) {
	assert.Equal(t, Sha256("123456"), "8d969eef6ecad3c29a3a629280e686cf0c3f5d5a86aff3ca12020c923adc6c92", "sha256不匹配")
}

func TestSha512(t *testing.T) {
	assert.Equal(t, Sha512("123456"), "ba3253876aed6bc22d4a6ff53d8406c6ad864195ed144ab5c87621b6c233b548baeae6956df346ec8c17f5ea10f35ee3cbc514797ed7ddd3145464e2a0bab413", "sha512不匹配")
}

func TestAesEncrypt(t *testing.T) {
	encrypt, _ := AesEncrypt("123456", "1234567890123456")
	assert.Equal(t, encrypt, "1jdzWuniG6UMtoa3T6uNLA==", "aes加密不匹配")
}

func TestAesDecrypt(t *testing.T) {
	decrypt, _ := AesDecrypt("1jdzWuniG6UMtoa3T6uNLA==", "1234567890123456")
	assert.Equal(t, decrypt, "123456", "aes解密不匹配")
}
