package responses_test

import (
	"encoding/json"
	"sync"

	"../responses"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Response Builder Tests 1", func() {

	Describe("When building a Response of type \"Proxy\"", func() {

		var (
			actualResponseAsMap map[string]interface{}

			once sync.Once
		)

		BeforeEach(func() {
			once.Do(func() {
				actualResponse := responses.Proxy().To("http://example.com").Mode("proxyOnce").Build()

				jsonBytes, _ := json.Marshal(actualResponse)
				actualResponseAsMap = map[string]interface{}{}
				json.Unmarshal(jsonBytes, &actualResponseAsMap)
			})
		})

		It("should create an \"Proxy\" response", func() {
			Expect(actualResponseAsMap).To(HaveKey("proxy"))
		})

		It("should create a Response with the correct To", func() {
			proxyResponse := actualResponseAsMap["proxy"]

			Expect(proxyResponse).To(HaveKeyWithValue("to", "http://example.com"))
		})

		It("should create a Response with Mode", func() {
			proxyResponse := actualResponseAsMap["proxy"]
			Expect(proxyResponse).To(HaveKeyWithValue("mode", "proxyOnce"))
		})
  })
})
