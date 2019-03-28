package dsig

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"hash"
	"strings"
)

func VerifySignature(key *rsa.PublicKey, toSign, signed []byte) error {
	h := sha256.New()
	h.Write(toSign)

	d := h.Sum(nil)
	err := rsa.VerifyPKCS1v15(key, crypto.SHA256, d, signed)
	if err != nil {
		return err
	}
	return nil
}

// CreateSignature will create a  digital signature for the data
// with the private key provided
func CreateSignature(data string, key []byte, algorithm string) (string, error) {
	block, _ := pem.Decode(key)
	if block == nil {
		return "", fmt.Errorf("invalid private key")
	}
	if block.Type != "RSA PRIVATE KEY" {
		return "", fmt.Errorf("no private key found")
	}
	pk, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}
	d, err := createHash([]byte(data), algorithm)
	if err != nil {
		return "", err
	}

	sig, err := rsa.SignPKCS1v15(rand.Reader, pk, crypto.SHA256, d)
	if err != nil {
		return "", err
	}
	return string(sig), nil
}

func createHash(data []byte, algo string) ([]byte, error) {
	var h hash.Hash
	switch algo {
	case "sha256":
		h = sha256.New()
		break
	case "sha1":
		h = sha1.New()
		break
	case "sha526":
		h = sha512.New()
		break
	}
	_, err := h.Write([]byte(data))
	if err != nil {
		return nil, err
	}
	return h.Sum(nil), nil
}

func clean(s string) string {
	var rs string
	for _, j := range strings.TrimSpace(s) {
		if string(j) != " " && string(j) != "\n" {
			rs += string(j)
		}

	}
	return rs
}
