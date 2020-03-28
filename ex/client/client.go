package client

import (
	"context"
	"crypto/aes"
	"crypto/cipher"

	"github.com/pkg/errors"
	"github.com/riyadennis/aes-encryption/data/models"
	"github.com/riyadennis/aes-encryption/ex"
	"github.com/riyadennis/aes-encryption/ex/api"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

type Client interface {
	Store(ctx context.Context, in *api.DataRequest,
		opts ...grpc.CallOption) (*api.DataResponse, error)

	// Retrieve accepts an id and an AES key, and requests that the
	// encryption-server retrieves the original (decrypted) bytes stored
	// with the provided id
	Retrieve(id, aesKey []byte) (payload []byte, err error)
}

type AesClient struct {
	Id   string `json:"id"`
	Key  string `json:"key;omitempty"`
	Data string `json:"data;omitempty"`
}

func (ac *AesClient) DataRequest() *api.DataRequest {
	return &api.DataRequest{
		Data: &api.Data{
			ToEncrypt:    ac.Data,
			EncryptionId: ac.Id,
		},
	}
}

func (ac *AesClient) Retrieve() (payload []byte, err error) {
	config, err := ex.GetConfig(ex.DefaultConfigPath)
	if err != nil {
		logrus.Errorf("unable to open config :: %v", err)
		return nil, err
	}
	data, err := models.GetPayLoad(ac.Id, config.Encrypter.Db)
	if err != nil {
		//already logged
		return nil, err
	}
	decryptedText, err := decrypt([]byte(data.EncryptedText), []byte(ac.Key))
	if err != nil {
		logrus.Errorf("decryption failed :: %v", err)
		return nil, err
	}
	return decryptedText, nil
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
