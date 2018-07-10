package handlers

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"encoding/json"
	"github.com/aes-encryption/client"
	"github.com/aes-encryption/middleware"
	"fmt"
)

type Input struct {
	Data string `json:"data"`
	Id   string `json:"id"`
}

func StoreDataHandler(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	var response *ApiResponse
	if req.Body == nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	requestBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	i := &Input{}
	err = json.Unmarshal(requestBody, i)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	config, err := middleware.GetConfigFromContext(req.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	ac := client.AesClient{Config: config}
	key, err := ac.Store([]byte(i.Id), []byte(i.Data))
	if err != nil {
		msg := fmt.Sprintf("Unable to store the data got error %s", err.Error())
		response = createResponse(msg, "Error", http.StatusBadRequest)
		jsonResponseDecorator(response, w)
		return
	}
	response = createResponse(string(key), "Success", http.StatusOK)
	jsonResponseDecorator(response, w)
}
