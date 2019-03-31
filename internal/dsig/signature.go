package dsig

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"strings"
)

// VerifySignature checkes signed data and the hash
func VerifySignature(
	key *rsa.PublicKey,
	toSign, signed []byte,
	hash crypto.Hash,
) error {
	h := hash.New()
	h.Write(toSign)

	d := h.Sum(nil)
	err := rsa.VerifyPKCS1v15(key, hash, d, signed)
	if err != nil {
		return err
	}
	return nil
}

// CreateSignature will create a  digital signature for the data
// with the private key provided
func CreateSignature(data string, key []byte, hash crypto.Hash) (string, error) {
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
	h := hash.New()
	h.Write([]byte(data))
	d := h.Sum(nil)

	sig, err := rsa.SignPKCS1v15(rand.Reader, pk, hash, d)
	if err != nil {
		return "", err
	}
	return string(sig), nil
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
