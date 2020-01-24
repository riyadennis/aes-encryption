package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/riyadennis/aes-encryption/ex"
	"github.com/riyadennis/aes-encryption/ex/api"

	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
)

// Run initialise handler
func Run(config *ex.Config) {
	route := httprouter.New()
	addr := fmt.Sprintf(":%d", config.Encrypter.Port)
	route.POST("/store", StoreDataHandler)
	route.GET("/get", GetDataHandler)
	fmt.Printf("Listenning to port %s \n", addr)
	logrus.Fatal(http.ListenAndServe(addr, ex.ConfigMiddleWare(route, config)))
}

func jsonResponseDecorator(response *api.DataResponse, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), int(response.GetHttpStatus()))
		return
	}
}
