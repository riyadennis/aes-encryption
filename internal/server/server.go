package server

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"fmt"
	"io"
	mathRand "math/rand"
	"net"
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/riyadennis/aes-encryption/data/models"
	"github.com/riyadennis/aes-encryption/ex"
	"github.com/riyadennis/aes-encryption/ex/api"
	"google.golang.org/grpc"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

type DataServiceServer struct {
	HttpStatus    int32
	EncryptionKey string
	Status        string
}

func (ds *DataServiceServer) Store(ctx context.Context, dr *api.DataRequest) (*api.DataResponse, error) {
	if dr.Data == nil {
		return nil, errors.New("invalid request")
	}
	key := randSeq(16)
	id := randSeq(32)

	encryptedText, err := encrypt(dr.Data.Message, id)
	if err != nil {
		return nil, err
	}
	cnf, err := ex.GetConfig(ex.DefaultConfigPath)
	if err != nil {
		return nil, err
	}
	err = models.SavePayload(id, key, encryptedText, cnf.Encrypter.Db)
	if err != nil {
		return nil, err
	}

	return &api.DataResponse{
		HttpStatus:    http.StatusOK,
		EncryptionKey: key,
		EncryptionId:  id,
	}, nil
}

func (ds *DataServiceServer) Retrieve(ctx context.Context, req *api.RetrieveRequest) (*api.RetrieveResponse, error) {
	config, err := ex.GetConfig(ex.DefaultConfigPath)
	if err != nil {
		logrus.Errorf("unable to open config :: %v", err)
		return nil, err
	}
	data, err := models.GetPayLoad(req.EncryptionId, config.Encrypter.Db)
	if err != nil {
		//already logged
		return nil, err
	}
	decryptedText, err := decrypt([]byte(data.EncryptedText), []byte(req.EncryptionKey))
	if err != nil {
		logrus.Errorf("decryption failed :: %v", err)
		return nil, err
	}
	return &api.RetrieveResponse{
		Data: &api.Data{Message: string(decryptedText)},
	}, nil
}

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[mathRand.Intn(len(letters))]
	}
	return string(b)
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
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return nil, err
	}
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

func Run() {
	addr := "0.0.0.0:5051"
	fmt.Printf("Listenning to port %s \n", addr)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	server := grpc.NewServer()
	s := &DataServiceServer{}
	api.RegisterDataServiceServer(server, s)
	if err = server.Serve(listener); err != nil {
		panic(err)
	}
}
