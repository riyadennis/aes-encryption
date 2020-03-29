package internal

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"

	"github.com/riyadennis/aes-encryption/internal/db"

	// TODO need to replace this with crypto/rand
	mathRand "math/rand"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func NewData(title, content string) (*db.Data, error) {
	key := randSeq(32)
	encryptedText, err := encrypt(content, key)
	if err != nil {
		return nil, err
	}
	return &db.Data{
		Key:     key,
		Title:   title,
		Content: string(encryptedText),
	}, nil
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

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[mathRand.Intn(len(letters))]
	}
	return string(b)
}
