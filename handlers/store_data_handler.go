package handlers

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"encoding/json"
	"github.com/aes-encryption/client"
	"github.com/aes-encryption/middleware"
)

type Input struct {
	Data string `json:"data"`
	Id   string `json:"id"`
}

func StoreDataHandler(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
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
	key, err := client.Store([]byte(i.Id), []byte(i.Data), config)
	w.Write([]byte(key))
}
