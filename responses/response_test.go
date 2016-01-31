package responses_test

import (
	"net/http"
	"sync"

	"github.com/durmaze/gobank/responses"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Response Builder Tests", func() {

	Describe("When building a Response of type \"Is\"", func() {

		var (
			actualResponse responses.Response

			expectedStatusCode  = http.StatusOK
			expectedHeader      = "Content-Type"
			expectedHeaderValue = "application/json"
			expectedBody        = "{ \"greeting\": \"Hello GoBank\" }"

			once sync.Once
		)

		BeforeEach(func() {
			once.Do(func() {
				actualResponse = responses.NewResponseBuilder().IsResponse().StatusCode(expectedStatusCode).Header(expectedHeader, expectedHeaderValue).Body(expectedBody).Build()
			})
		})

		It("should create a Response with the correct StatusCode", func() {
			Expect(actualResponse.Is.StatusCode).To(Equal(expectedStatusCode))
		})

		It("should create a Response with the correct Header", func() {
			Expect(actualResponse.Is.Headers[expectedHeader]).To(Equal(expectedHeaderValue))
		})

		It("should create a Response with the correct Body", func() {
			Expect(actualResponse.Is.Body).To(Equal(expectedBody))
		})
	})

	XDescribe("When building a Response of type \"Proxy\"", func() {
		// To be implemented
	})

})
