package handlers_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"net/http"
	"strings"
	"github.com/aes-encryption/middleware"
	"github.com/aes-encryption/handlers"
)

var _ = Describe("GetDataHandler", func() {

	BeforeEach(func() {
		config, err := middleware.GetConfig("../config_test.yaml")
		Expect(err).To(BeNil())
		go handlers.Run(config)

	})

	Context("get bad request for a bad post", func() {
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

})
