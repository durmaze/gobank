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

		It("should not have 'behavior' field in marshalled json", func() {
			Expect(actualResponseAsMap).NotTo(HaveKey("_behaviors"))
		})
	})

	Describe("When building a Response of type \"Is\" with behavior", func() {

		var (
			actualResponseAsMap map[string]interface{}
			once                sync.Once
		)

		const (
			expectedWaitTime = 500
			expectedBody     = "{ \"greeting\": \"Hello GoBank\" }"
		)

		BeforeEach(func() {
			once.Do(func() {
				actualResponse := responses.Is().Body(expectedBody).Wait(expectedWaitTime).Build()

				jsonBytes, _ := json.Marshal(actualResponse)
				actualResponseAsMap = map[string]interface{}{}
				json.Unmarshal(jsonBytes, &actualResponseAsMap)
			})
		})

		It("should create an \"Is\" response", func() {
			Expect(actualResponseAsMap).To(HaveKey("is"))
		})

		It("should create a Response with the correct Body", func() {
			isResponse := actualResponseAsMap["is"]

			Expect(isResponse).To(HaveKeyWithValue("body", expectedBody))
		})

		It("should have 'behavior' field in marshalled json", func() {
			Expect(actualResponseAsMap).To(HaveKey("_behaviors"))
		})

		It("should have 'wait' field in as 'behavior' in marshalled json", func() {
			behavior := actualResponseAsMap["_behaviors"]
			Expect(behavior).To(HaveKeyWithValue("wait", float64(expectedWaitTime)))
		})

		It("should not have 'decorate' field in as 'behavior' in marshalled json", func() {
			behavior := actualResponseAsMap["_behaviors"]
			Expect(behavior).NotTo(HaveKey("decorate"))
		})
	})

	Describe("When building a Response of type \"Is\" with behavior", func() {

		var (
			actualResponseAsMap map[string]interface{}
			once                sync.Once
		)

		const (
			expectedWaitTime = 500
			expectedBody     = "{ \"greeting\": \"Hello GoBank\" }"
		)

		BeforeEach(func() {
			once.Do(func() {
				actualResponse := responses.Is().Body(expectedBody).Decorate("function(){}").Build()

				jsonBytes, _ := json.Marshal(actualResponse)
				actualResponseAsMap = map[string]interface{}{}
				json.Unmarshal(jsonBytes, &actualResponseAsMap)
			})
		})

		It("should create an \"Is\" response", func() {
			Expect(actualResponseAsMap).To(HaveKey("is"))
		})

		It("should create a Response with the correct Body", func() {
			isResponse := actualResponseAsMap["is"]

			Expect(isResponse).To(HaveKeyWithValue("body", expectedBody))
		})

		It("should have 'behavior' field in marshalled json", func() {
			Expect(actualResponseAsMap).To(HaveKey("_behaviors"))
		})

		It("should not have 'wait' field in as 'behavior' in marshalled json", func() {
			behavior := actualResponseAsMap["_behaviors"]
			Expect(behavior).NotTo(HaveKey("wait"))
		})

		It("should have 'decorate' field in as 'behavior' in marshalled json", func() {
			behavior := actualResponseAsMap["_behaviors"]
			Expect(behavior).To(HaveKey("decorate"))
		})
	})

	Describe("When building a Response of type \"Is\" with behavior", func() {

		var (
			actualResponseAsMap map[string]interface{}
			once                sync.Once
		)

		const (
			expectedBody = "{ \"greeting\": \"Hello GoBank\" }"
		)

		BeforeEach(func() {
			once.Do(func() {
				actualResponse := responses.Is().Body(expectedBody).Wait(0).Build()

				jsonBytes, _ := json.Marshal(actualResponse)
				actualResponseAsMap = map[string]interface{}{}
				json.Unmarshal(jsonBytes, &actualResponseAsMap)
			})
		})

		It("should create an \"Is\" response", func() {
			Expect(actualResponseAsMap).To(HaveKey("is"))
		})

		It("should create a Response with the correct Body", func() {
			isResponse := actualResponseAsMap["is"]

			Expect(isResponse).To(HaveKeyWithValue("body", expectedBody))
		})

		It("should not have 'behavior' field in marshalled json", func() {
			Expect(actualResponseAsMap).NotTo(HaveKey("_behavior"))
		})

	})

})
