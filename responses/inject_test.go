package responses_test

import (
	"encoding/json"
	"sync"

	"github.com/durmaze/gobank/responses"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Response Builder Tests", func() {

	Describe("When building a Response of type \"Is\"", func() {

		var (
			actualResponseAsMap map[string]interface{}

			expectedFunction = "function (request, state, logger) { logger.info('origin called'); return { headers: { 'Content-Type': 'application/json' }, body: JSON.stringify({ count: state.requests })};}"

			once sync.Once
		)

		BeforeEach(func() {
			once.Do(func() {
				actualResponse := responses.Inject().
					Fn(expectedFunction).
					Build()

				jsonBytes, _ := json.Marshal(actualResponse)
				actualResponseAsMap = map[string]interface{}{}
				json.Unmarshal(jsonBytes, &actualResponseAsMap)
			})
		})

		It("should create an \"inject\" response", func() {
			Expect(actualResponseAsMap).To(HaveKeyWithValue("inject", expectedFunction))
		})

	})

})
