package server

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"errors"
	"net/http"

	"github.com/riyadennis/aes-encryption/internal"

	"github.com/riyadennis/aes-encryption/ex/api"
	"github.com/sirupsen/logrus"
)

type DataServiceServer struct {
	HttpStatus    int32
	EncryptionKey string
	Status        string
}

func (ds *DataServiceServer) Store(ctx context.Context,
	dr *api.DataRequest) (*api.DataResponse, error) {
	if dr.Data == nil {
		return nil, errors.New("invalid request")
	}
	data, err := internal.NewData("static title", dr.Data.Message)
	if err != nil {
		return nil, err
	}
	re, err := store.Insert(ctx, data)
	if err != nil {
		return nil, err
	}
	return &api.DataResponse{
		HttpStatus:    http.StatusOK,
		EncryptionKey: re.EncryptionKey,
		EncryptionId:  re.EncryptionId,
	}, nil
}

func (ds *DataServiceServer) Retrieve(ctx context.Context, req *api.RetrieveRequest) (*api.RetrieveResponse, error) {
	data, err := store.RetrieveContent(ctx, req.EncryptionId)
	if err != nil {
		logrus.Errorf("read from db failed :: %v", err)
		return nil, err
	}
	decryptedText, err := decrypt([]byte(data.Content), []byte(req.EncryptionKey))
	if err != nil {
		logrus.Errorf("decryption failed :: %v", err)
		return nil, err
	}
	return &api.RetrieveResponse{
		Data: &api.Data{Message: string(decryptedText)},
	}, nil
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
