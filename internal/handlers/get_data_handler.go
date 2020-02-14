package handlers

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/riyadennis/aes-encryption/ex/api"
	"github.com/riyadennis/aes-encryption/ex/client"
)

// GetDataHandler handles rest call to fetch data from the db
func GetDataHandler(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	key := req.Header.Get("key")
	id := req.Header.Get("id")
	if key == "" || id == "" {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	ac := &client.AesClient{
		Id:  id,
		Key: key,
	}
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
