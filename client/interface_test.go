package client

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestRandSeq(t *testing.T) {
	size := 16
	key := RandSeq(size)
	assert.Equal(t, 16, len(key))

	size1 := 12
	key1 := RandSeq(size1)
	assert.Equal(t, 12, len(key1))
}
func TestEncrypt(t *testing.T) {
	key := RandSeq(16)
	textToEncrypt := "Plain text"
	encryptedText, err := Encrypt(textToEncrypt, key)
	assert.NoError(t, err)
	assert.NotEmpty(t, encryptedText)
}

func TestDecrypt(t *testing.T) {
	key := RandSeq(16)
	textToEncrypt := "Plain text"
	encryptedText, err := Encrypt(textToEncrypt, key)
	assert.NoError(t, err)
	decryptedText, err  := Decrypt(encryptedText, []byte(key))
	assert.NoError(t, err)
	assert.Equal(t, string(decryptedText), textToEncrypt)
}