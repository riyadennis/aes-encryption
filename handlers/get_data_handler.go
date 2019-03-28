package handlers

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/riyadennis/aes-encryption/middleware"
	"github.com/riyadennis/aes-encryption/client"
	"fmt"
)


func GetDataHandler(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	var response *ApiResponse
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
		response = createResponse("Unable to retrieve data", "Error", http.StatusBadRequest)
		jsonResponseDecorator(response, w)
		return
	}
	cipherText := fmt.Sprintf("cipher_text: %s", string(data))
	response = createResponse(cipherText, "Success", http.StatusOK)
	jsonResponseDecorator(response, w)
}
