package responses_test

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/durmaze/gobank/responses"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Response Builder Tests", func() {

	Describe("When building a Response of type \"Is\"", func() {

		var (
			actualResponseAsMap map[string]interface{}

			expectedStatusCode  = http.StatusOK
			expectedHeader      = "Content-Type"
			expectedHeaderValue = "application/json"
			expectedBody        = "{ \"greeting\": \"Hello GoBank\" }"

			once sync.Once
		)

		BeforeEach(func() {
			once.Do(func() {
				actualResponse := responses.Is().StatusCode(expectedStatusCode).Header(expectedHeader, expectedHeaderValue).Body(expectedBody).Build()

				jsonBytes, _ := json.Marshal(actualResponse)
				actualResponseAsMap = map[string]interface{}{}
				json.Unmarshal(jsonBytes, &actualResponseAsMap)
			})
		})

		It("should create an \"Is\" response", func() {
			Expect(actualResponseAsMap).To(HaveKey("is"))
		})

		It("should create a Response with the correct StatusCode", func() {
			isResponse := actualResponseAsMap["is"]

			// golang converts numbers to float64 when unmarshalling json to map[string]interface{}
			Expect(isResponse).To(HaveKeyWithValue("statusCode", float64(expectedStatusCode)))
		})

		It("should create a Response with Headers", func() {
			isResponse := actualResponseAsMap["is"]
			Expect(isResponse).To(HaveKey("headers"))
		})
		It("should create a Response with Headers", func() {
			isResponse := actualResponseAsMap["is"].(map[string]interface{})
			headersMap := isResponse["headers"]

			Expect(headersMap).To(HaveKeyWithValue(expectedHeader, expectedHeaderValue))
		})

		It("should create a Response with the correct Body", func() {
			isResponse := actualResponseAsMap["is"]

			Expect(isResponse).To(HaveKeyWithValue("body", expectedBody))
		})
	})

})
