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
uq3RY9isrlVDrcUC979koHH/9+8LxCUOj3ukH3RVJsHsEYtJm+FA1JaWQSrVFhtb
03fen4SY/w3AoHui56BMuW2OQpAXDjm637ooYcYLET0GTofOH5EbW/CU/OYdWMf/
Vct4fPGq5oTekspkynYvGYXU/A5iYKmsVKCC/ZwRmNY=`
	messageToSign = "data"
)

func TestCreateSignatureWithHashSha256(t *testing.T) {
	t.Parallel()
	key, err := ioutil.ReadFile("server.key")
	if err != nil {
		t.Errorf("failed :: %v", err)
	}
	sig, err := CreateSignature(messageToSign, key, crypto.SHA256)
	if err != nil {
		t.Errorf("failed :: %v", err)
	}
	encodedSv := base64.StdEncoding.EncodeToString([]byte(sig))
	fmt.Printf("%v", encodedSv)
	if encodedSv != clean(expectedSha256) {
		t.Errorf("mismatch")
	}
}

func TestCreateSignatureWithHashSha1(t *testing.T) {
	t.Parallel()
	expectedSha1 := `
xyPRuW0KXqLO47PGwMv1mflRDxMV7P+cnPHuNb4JaHajxus9B0U5Ai3wQ4SKWbXZoAe5
XcblXhHCVgo1Jphnjb29zAocuDjj5PnokJ16l9UzJjaGRNtNYs9E6Rvn6DJE6nt0IESi
Mg8yePSnoLlSyVlZEod3RvAfH9tn9sCkK94= `

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
	hashAlgo := crypto.SHA1
	key, err := ioutil.ReadFile("server.key")
	if err != nil {
		t.Errorf("failed :: %v", err)
	}
	sig, err := CreateSignature(messageToSign, key, hashAlgo)
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
	err = VerifySignature(publicKey, []byte(messageToSign), []byte(sig), hashAlgo)
	if err != nil {
		t.Errorf("failed :: %v", err)
	}
}
