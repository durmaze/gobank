package predicates_test

import (
	"encoding/json"
	"sync"

	"github.com/durmaze/gobank/predicates"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Equals Predicate Builder Tests", func() {

	Describe("When building a Predicate of type \"Equals\"", func() {

		var (
			actualPredicateAsMap map[string]interface{}

			expectedPath            = "/testing/path"
			expectedMethod          = "GET"
			expectedHeader          = "Content-Type"
			expectedHeaderValue     = "application/json"
			expectedBody            = "{ \"greeting\": \"Hello GoBank\" }"
			expectedQueryParam      = "qp"
			expectedQueryParamValue = "testValue"

			once sync.Once
		)

		BeforeEach(func() {
			once.Do(func() {
				actualPredicate := predicates.Equals().
					Path(expectedPath).
					Method(expectedMethod).
					Header(expectedHeader, expectedHeaderValue).
					Query(expectedQueryParam, expectedQueryParamValue).
					Body(expectedBody).
					Build()

				jsonBytes, _ := json.Marshal(actualPredicate)
				actualPredicateAsMap = map[string]interface{}{}
				json.Unmarshal(jsonBytes, &actualPredicateAsMap)

				// marshalResult := NewMarshaller(actualPredicate).ToJson()
				// actualPredicateAsMap = marshalResult.ToMap()
			})
		})

		It("should create a \"Equals\" Predicate", func() {
			Expect(actualPredicateAsMap).To(HaveKey("equals"))
		})

		It("should create a Predicate with the correct Path", func() {
			equalsPredicate := actualPredicateAsMap["equals"]

			Expect(equalsPredicate).To(HaveKeyWithValue("path", expectedPath))
		})

		It("should create a Predicate with the correct Method", func() {
			equalsPredicate := actualPredicateAsMap["equals"]

			Expect(equalsPredicate).To(HaveKeyWithValue("method", expectedMethod))
		})

		It("should create a Predicate with Headers", func() {
			equalsPredicate := actualPredicateAsMap["equals"]

			Expect(equalsPredicate).To(HaveKey("headers"))
		})

		It("should create a Predicate with the correct Header", func() {
			equalsPredicate := actualPredicateAsMap["equals"].(map[string]interface{})

			headersMap := equalsPredicate["headers"].(map[string]interface{})
			headerValue := headersMap[expectedHeader]

			Expect(headerValue).To(Equal(expectedHeaderValue))
		})

		It("should create a Predicate with Query parameters", func() {
			equalsPredicate := actualPredicateAsMap["equals"]

			Expect(equalsPredicate).To(HaveKey("query"))
		})

		It("should create a Predicate with the correct Query parameter", func() {
			equalsPredicate := actualPredicateAsMap["equals"].(map[string]interface{})

			queryParametersMap := equalsPredicate["query"].(map[string]interface{})

			Expect(queryParametersMap).To(HaveKeyWithValue(expectedQueryParam, expectedQueryParamValue))
		})

		It("should create a Predicate with the correct Body", func() {
			equalsPredicate := actualPredicateAsMap["equals"]

			Expect(equalsPredicate).To(HaveKeyWithValue("body", expectedBody))
		})

	})

})
