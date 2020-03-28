package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/riyadennis/aes-encryption/data/models"
	"github.com/riyadennis/aes-encryption/ex"
	"github.com/riyadennis/aes-encryption/ex/api"
	"github.com/riyadennis/aes-encryption/ex/client"
)

func StoreDataHandler(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	var resp *api.DataResponse
	if req.Body == nil {
		jsonResponseDecorator(&api.DataResponse{
			HttpStatus: http.StatusBadRequest,
			Status:     fmt.Sprintf("invalid body"),
		}, w)
		return
	}
	requestBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		jsonResponseDecorator(&api.DataResponse{
			HttpStatus: http.StatusBadRequest,
			Status:     fmt.Sprintf("%v", err),
		}, w)
		return
	}
	ac := &client.AesClient{}
	err = json.Unmarshal(requestBody, ac)
	if err != nil {
		jsonResponseDecorator(&api.DataResponse{
			HttpStatus: http.StatusBadRequest,
			Status:     fmt.Sprintf("%v", err),
		}, w)
		return
	}
	cnf, err := ex.GetConfig(ex.DefaultConfigPath)
	if err != nil {
		jsonResponseDecorator(&api.DataResponse{
			HttpStatus: http.StatusBadRequest,
			Status:     fmt.Sprintf("%v", err),
		}, w)
		return
	}
	err = models.SavePayload(ac.Id, ac.Key,
		[]byte(ac.Data), cnf.Encrypter.Db)
	if err != nil {
		jsonResponseDecorator(&api.DataResponse{
			HttpStatus: http.StatusInternalServerError,
			Status:     fmt.Sprintf("%v", err),
		}, w)
		return
	}
	resp = &api.DataResponse{
		HttpStatus:    http.StatusOK,
		Status:        "Success",
		EncryptionKey: ac.Key,
	}
	jsonResponseDecorator(resp, w)
}
