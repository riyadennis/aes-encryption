package server

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	mathRand "math/rand"
	"net"
	"net/http"

	"github.com/riyadennis/aes-encryption/data/models"
	"github.com/riyadennis/aes-encryption/ex"
	"github.com/riyadennis/aes-encryption/ex/api"
	"google.golang.org/grpc"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

type AesServer struct {
	HttpStatus    int32
	EncryptionKey string
	Status        string
}

func (ae *AesServer) Store(ctx context.Context, dr *api.DataRequest) (*api.DataResponse, error) {
	key := randSeq(16)
	encryptedText, err := encrypt(dr.Data.ToEncrypt, dr.Data.EncryptionId)
	if err != nil {
		return nil, err
	}
	cnf, err := ex.GetConfig(ex.DefaultConfigPath)
	if err != nil {
		return nil, err
	}
	err = models.SavePayload(dr.Data.EncryptionId, key, encryptedText, cnf.Encrypter.Db)
	if err != nil {
		return nil, err
	}

	return &api.DataResponse{
		HttpStatus:    http.StatusOK,
		EncryptionKey: key,
		Status:        "Success",
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

func Run() {
	listener, err := net.Listen("tcp", "0.0.0.0:5051")
	if err != nil {
		panic(err)
	}
	server := grpc.NewServer()
	api.RegisterAesServer(server, &AesServer{})
	if err = server.Serve(listener); err != nil {
		panic(err)
	}
}
