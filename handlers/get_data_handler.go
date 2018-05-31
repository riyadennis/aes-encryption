package handlers

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/aes-encryption/models"
	"github.com/aes-encryption/middleware"
	"github.com/aes-encryption/client"
	"fmt"
)

func GetDataHandler(w http.ResponseWriter, req *http.Request, _ httprouter.Params){
	key := req.Header.Get("key")
	id := req.Header.Get("id")
	config, err := middleware.GetConfigFromContext(req.Context())
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	data, err := models.GetPayLoad(id, config.Encrypter.Db)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	decryptedText, err := client.Decrypt([]byte(data.EncryptedText), []byte(key))
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	cipherText := fmt.Sprintf("cipher_text: %s", decryptedText)
	response := createResponse(cipherText, "Success", http.StatusOK)
	jsonResponseDecorator(response, w)
}
