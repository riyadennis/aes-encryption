package handlers

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/aes-encryption/middleware"
	"github.com/aes-encryption/client"
	"fmt"
)

func GetDataHandler(w http.ResponseWriter, req *http.Request, _ httprouter.Params){
	key := req.Header.Get("key")
	id := req.Header.Get("id")
	if key == ""{
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	config, err := middleware.GetConfigFromContext(req.Context())
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	ac := client.AesClient{Config: config}
	data, err := ac.Retrieve([]byte(id), []byte(key))
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	cipherText := fmt.Sprintf("cipher_text: %s", string(data))
	response := createResponse(cipherText, "Success", http.StatusOK)
	jsonResponseDecorator(response, w)
}
