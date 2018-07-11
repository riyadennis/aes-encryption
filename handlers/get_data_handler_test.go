package handlers_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"net/http"
	"strings"
	"github.com/aes-encryption/middleware"
	"github.com/aes-encryption/handlers"
	"fmt"
)

var _ = Describe("GetDataHandler", func() {

	BeforeSuite(func() {
		config, err := middleware.GetConfig("../config_test.yaml")
		Expect(err).To(BeNil())
		go handlers.Run(config)

	})

	Context("get bad request for a post", func() {
		It("Should give me a valid text", func() {
			data := `{
	"data": "functions useful for working with such data",
	"id": "eeee22234dsfsdfdsfddd"
}`
			message := strings.NewReader(data)
			resp, err := http.Post("http://localhost:8081/store", "application/json", message)
			Expect(err).To(BeNil())
			Expect(resp.StatusCode).To(Equal(http.StatusBadRequest))
		})
	})

	Context("get status bad request no header get", func() {
		It("Should give me a valid text", func() {
			resp, err := http.Get("http://localhost:8081/get")
			Expect(err).To(BeNil())
			Expect(resp.StatusCode).To(Equal(http.StatusBadRequest))
		})
	})

	Context("get status bad request no id get", func() {
		It("Should give error", func() {
			req, err := http.NewRequest("GET", "http://localhost:8081/get", nil)
			req.Header.Add("key", "test key")
			client := http.Client{}
			resp, err := client.Do(req)
			Expect(err).To(BeNil())
			Expect(resp.StatusCode).To(Equal(http.StatusBadRequest))
		})
	})

	Context("get status bad request for id and key that is not in db", func() {
		It("Should give me error", func() {
			var b []byte
			req, err := http.NewRequest("GET", "http://localhost:8081/get", nil)
			req.Header.Add("key", "test key")
			req.Header.Add("id", "test id")
			client := http.Client{}
			resp, err := client.Do(req)
			Expect(err).To(BeNil())
			resp.Body.Read(b)
			fmt.Printf("%#v", string(b))
			Expect(resp.StatusCode).To(Equal(http.StatusBadRequest))
		})
	})

})
