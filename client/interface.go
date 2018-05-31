package client

import (
	"crypto/aes"
	"crypto/cipher"
	"io"
	"crypto/rand"
	"github.com/pkg/errors"
	mathRand "math/rand"
	"github.com/aes-encryption/models"
	"github.com/aes-encryption/middleware"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

type Client interface {
	// Store accepts an id and a payload in bytes and requests that the
	// encryption-server stores them in its data store
	Store(id, payload []byte, config *middleware.Config) (aesKey []byte, err error)

	// Retrieve accepts an id and an AES key, and requests that the
	// encryption-server retrieves the original (decrypted) bytes stored
	// with the provided id
	Retrieve(id, aesKey []byte) (payload []byte, err error)
}

func Store(id, payload []byte, config *middleware.Config) (aesKey []byte, err error) {
	if id == nil {
		return nil, errors.New("Invalid input")
	}

	key := RandSeq(16)
	encryptedText, err := Encrypt(string(payload), key)
	err = models.SavePayload(string(id), key, encryptedText, config.Encrypter.Db)
	if err != nil {
		return nil, err
	}
	return []byte(key), nil
}

func Retrieve(id, aesKey []byte) (payload []byte, err error) {
	conf, err := middleware.GetConfig("../config.yaml")
	if err != nil {
		return nil, err
	}
	data, err := models.GetPayLoad(string(id), conf.Encrypter.Db)
	decryptedText, err := Decrypt([]byte(data.EncryptedText), aesKey)
	if err != nil {
		return nil, err
	}
	return decryptedText, nil
}

func Encrypt(plainText, key string) ([]byte, error) {
	c, err := aes.NewCipher([]byte(key))
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}
	nonce := make([]byte, gcm.NonceSize())
	io.ReadFull(rand.Reader, nonce)
	return gcm.Seal(nonce, nonce, []byte(plainText), nil), nil
}
func Decrypt(encryptedText, key []byte) ([]byte, error) {
	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}
	nonceSize := gcm.NonceSize()
	if len(encryptedText) < nonceSize {
		return nil, errors.New("Error encrypted text is too small")
	}
	nonce, decryptedText := encryptedText[:nonceSize], encryptedText[nonceSize:]
	return gcm.Open(nil, nonce, decryptedText, nil)
}

func RandSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[mathRand.Intn(len(letters))]
	}
	return string(b)
}
