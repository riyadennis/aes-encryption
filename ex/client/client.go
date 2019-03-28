package client

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	mathRand "math/rand"

	"github.com/pkg/errors"
	"github.com/riyadennis/aes-encryption/ex/api"
	"github.com/riyadennis/aes-encryption/internal"
	"github.com/riyadennis/aes-encryption/internal/models"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

type Client interface {
	Store(payLoad, encryptionId string) *api.DataRequest

	// Retrieve accepts an id and an AES key, and requests that the
	// encryption-server retrieves the original (decrypted) bytes stored
	// with the provided id
	Retrieve(id, aesKey []byte) (payload []byte, err error)
}

type AesClient struct{}

func (ac AesClient) DataRequest(payLoad, encryptionId string) *api.DataRequest {
	return &api.DataRequest{
		Data: &api.Data{
			ToEncrypt:    payLoad,
			EncryptionId: encryptionId,
		},
	}
}

func (ac AesClient) Retrieve(id, aesKey []byte) (payload []byte, err error) {
	config, err := internal.GetConfig(internal.DefaultConfigPath)
	if err != nil {
		return nil, err
	}
	data, err := models.GetPayLoad(string(id), config.Encrypter.Db)
	if err != nil {
		return nil, err
	}
	decryptedText, err := decrypt([]byte(data.EncryptedText), aesKey)
	if err != nil {
		return nil, err
	}
	return decryptedText, nil
}

func encrypt(plainText, key string) ([]byte, error) {
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
func decrypt(encryptedText, key []byte) ([]byte, error) {
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

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[mathRand.Intn(len(letters))]
	}
	return string(b)
}
