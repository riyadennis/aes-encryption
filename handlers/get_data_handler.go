package handlers

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/riyadennis/aes-encryption/ex/api"
	"github.com/riyadennis/aes-encryption/ex/client"
	"github.com/riyadennis/aes-encryption/middleware"
)

func GetDataHandler(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	key := req.Header.Get("key")
	id := req.Header.Get("id")
	if key == "" || id == "" {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	config, err := middleware.GetConfigFromContext(req.Context())
	if err != nil {
		http.Error(w, "unable to fetch config :: %v", http.StatusBadRequest)
		return
	}
	ac := client.AesClient{Config: config}
	data, err := ac.Retrieve([]byte(id), []byte(key))
	if err != nil {
		response := &api.DataResponse{HttpStatus: http.StatusInternalServerError}
		jsonResponseDecorator(response, w)
		return
	}
	cipherText := fmt.Sprintf("cipher_text: %s", string(data))
	response := &api.DataResponse{
		HttpStatus:    http.StatusOK,
		Status:        "Success",
		EncryptionKey: cipherText,
	}
	jsonResponseDecorator(response, w)
}
