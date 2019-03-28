package dsig

import (
	"crypto"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"testing"
)

var (
	expectedSha256 = `
flTCqX40Lf4vdEhWcqdPrC4AS6KLEqOlaoscUAyxqzb
1Hw8XmbsZF+BibhgXl1L+7kf0gevnQJ+MKDDfsazmW+
NOMiNusOTai/uwe5OnoSjRUZXORFCvJEGmPXM7mkp64
IFNxvoGzk6yq3fCrFcNHTP81cto27434reOU8LJsXo=`
	messageToSign = "data"
)

func TestCreateSignatureWithHashSha256(t *testing.T) {
	key, err := ioutil.ReadFile("server.key")
	if err != nil {
		t.Errorf("failed :: %v", err)
	}
	sig, err := CreateSignature(messageToSign, key, crypto.SHA256)
	if err != nil {
		t.Errorf("failed :: %v", err)
	}
	encodedSv := base64.StdEncoding.EncodeToString([]byte(sig))
	if encodedSv != clean(expectedSha256) {
		t.Errorf("mismatch")
	}
}

func TestCreateSignatureWithHashSha1(t *testing.T) {
	expectedSha1 := `
OSH1HartHf48R88LPod8BkB1atpJJrIAwrim0iO3Og+x
QIiCI3xazErcRZjOrZnJRTUfRYPcyXMPx46f+CFD2O3P
zyShTjwvCh5QwfRcOuvRmvlA5MfIG0nKxpBG71YSAujr
7WHmzk8KyhQ5nHuMsDJ84l3PjQA6BuvzVXTOE5o= `

	key, err := ioutil.ReadFile("server.key")
	if err != nil {
		t.Errorf("failed :: %v", err)
	}
	sig, err := CreateSignature(messageToSign, key, crypto.SHA1)
	if err != nil {
		t.Errorf("failed :: %v", err)
	}
	encodedSv := base64.StdEncoding.EncodeToString([]byte(sig))
	if encodedSv != clean(expectedSha1) {
		t.Errorf("mismatch")
	}
}

func TestSignatureVerification(t *testing.T) {
	hash := crypto.SHA1
	key, err := ioutil.ReadFile("key.pem")
	if err != nil {
		t.Errorf("failed :: %v", err)
	}
	sig, err := CreateSignature(messageToSign, key, hash)
	if err != nil {
		t.Errorf("failed :: %v", err)
	}

	encodedSv := base64.StdEncoding.EncodeToString([]byte(sig))
	fmt.Printf("%v \n", encodedSv)

	cer, err := ioutil.ReadFile("certificate.pem")
	if err != nil {
		t.Errorf("failed :: %v", err)
	}
	block, _ := pem.Decode(cer)
	if block == nil {
		t.Errorf("invalid certificate no PEM data found")
	}
	if block.Type != "CERTIFICATE" {
		t.Errorf("invalid certificate")
	}
	c, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		t.Error(err)
	}
	publicKey, ok := c.PublicKey.(*rsa.PublicKey)
	if !ok {
		t.Errorf("certificate's public key is not RSA")
	}
	err = VerifySignature(publicKey, []byte("data to be signed"), []byte(sig), hash)
	if err != nil {
		t.Errorf("failed :: %v", err)
	}
}
