package dsig

import (
	"encoding/base64"
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
	sig, err := CreateSignature(messageToSign, key, "sha256")
	if err != nil {
		t.Errorf("failed :: %v", err)
	}
	encodedSv := base64.StdEncoding.EncodeToString([]byte(sig))
	if encodedSv != clean(expectedSha256) {
		t.Errorf("mismatch")
	}
}

func TestCreateSignatureWithHashSha1(t *testing.T) {
	key, err := ioutil.ReadFile("server.key")
	if err != nil {
		t.Errorf("failed :: %v", err)
	}
	sig, err := CreateSignature(messageToSign, key, "sha526")
	if err != nil {
		t.Errorf("failed :: %v", err)
	}
	encodedSv := base64.StdEncoding.EncodeToString([]byte(sig))
	if encodedSv != clean(expectedSha256) {
		t.Errorf("mismatch")
	}
}

//func TestSignatureVerification(t *testing.T) {
//	sig, err := CreateSignature("data to be signed", []byte(testPrivateKey))
//	if err != nil {
//		t.Errorf("failed :: %v", err)
//	}
//
//	encodedSv := base64.StdEncoding.EncodeToString([]byte(sig))
//	fmt.Printf("%v", encodedSv)
//
//	cer, err := ioutil.ReadFile(testPrivateKey)
//	if err != nil {
//		t.Errorf("failed :: %v", err)
//	}
//	block, _ := pem.Decode(cer)
//	if block == nil {
//		t.Errorf("invalid certificate no PEM data found")
//	}
//	if block.Type != "CERTIFICATE" {
//		t.Errorf("invalid certificate")
//	}
//	c, err := x509.ParseCertificate(block.Bytes)
//	if err != nil {
//		t.Error(err)
//	}
//	key, ok := c.PublicKey.(*rsa.PublicKey)
//	if !ok {
//		t.Errorf("certificate's public key is not RSA")
//	}
//	err = VerifySignature(key, []byte("data to be signed"), []byte(sig))
//	if err != nil {
//		t.Errorf("failed :: %v", err)
//	}
//}
