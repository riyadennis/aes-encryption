package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/riyadennis/aes-encryption/ex/api"
	"github.com/riyadennis/aes-encryption/ex/client"
	"github.com/riyadennis/aes-encryption/internal/server"
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
	ac := client.AesClient{}

	re := ac.DataRequest(i.Data, i.Id)

	s := server.AesServer{
		HttpStatus:    http.StatusOK,
		EncryptionKey: re.Data.ToEncrypt,
		Status:        "Success",
	}
	resp, err := s.Store(context.Background(), re)
	if err != nil {
		resp = &api.DataResponse{
			HttpStatus: http.StatusInternalServerError,
			Status:     fmt.Sprintf("%v", err),
		}
	}
	jsonResponseDecorator(resp, w)
}
