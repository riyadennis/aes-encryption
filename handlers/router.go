package handlers

import (
	"github.com/julienschmidt/httprouter"
	"fmt"
	"encoding/json"
	"net/http"
	"github.com/aes-encryption/middleware"
	"github.com/sirupsen/logrus"
	"os"
)


type ApiResponse struct {
	Status int
	Detail string
	Title  string
}
func Run(config *middleware.Config) {
	route := httprouter.New()
	addr := fmt.Sprintf(":%d", config.Encrypter.Port)
	fmt.Printf("Listenning to port %s \n", addr)
	logrus.Fatal(http.ListenAndServe(addr, middleware.ConfigMiddleWare(route, config)))
}

func jsonResponseDecorator(response *ApiResponse, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), response.Status)
		return
	}
}
func createResponse(detail, title string, status int) *ApiResponse {
	return &ApiResponse{
		Status: status,
		Detail: detail,
		Title:  title,
	}
}
func GetLocalImageURL(config *middleware.Config, filename, fileType string) string{
	hostName, _ := os.Hostname()
	return fmt.Sprintf("Image URL: http://%s:%d/%s.%s",hostName, config.Encrypter.Port, filename, fileType)
}