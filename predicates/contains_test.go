package predicates_test

import (
	"encoding/json"
	"sync"

	"github.com/durmaze/gobank/predicates"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Contains Predicate Builder Tests", func() {

	Describe("When building a Predicate of type \"Contains\"", func() {

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
				actualPredicate := predicates.Contains().
					Path(expectedPath).
					Method(expectedMethod).
					Header(expectedHeader, expectedHeaderValue).
					Query(expectedQueryParam, expectedQueryParamValue).
					Body(expectedBody).
					Build()

				jsonBytes, _ := json.Marshal(actualPredicate)
				actualPredicateAsMap = map[string]interface{}{}
				json.Unmarshal(jsonBytes, &actualPredicateAsMap)
			})
		})

		It("should create a \"Contains\" Predicate", func() {
			Expect(actualPredicateAsMap).To(HaveKey("contains"))
		})

		It("should create a Predicate with the correct Path", func() {
			containsPredicate := actualPredicateAsMap["contains"]

			Expect(containsPredicate).To(HaveKeyWithValue("path", expectedPath))
		})

		It("should create a Predicate with the correct Method", func() {
			containsPredicate := actualPredicateAsMap["contains"]

			Expect(containsPredicate).To(HaveKeyWithValue("method", expectedMethod))
		})

		It("should create a Predicate with Headers", func() {
			containsPredicate := actualPredicateAsMap["contains"]

			Expect(containsPredicate).To(HaveKey("headers"))
		})

		It("should create a Predicate with the correct Header", func() {
			containsPredicate := actualPredicateAsMap["contains"].(map[string]interface{})

			headersMap := containsPredicate["headers"].(map[string]interface{})
			headerValue := headersMap[expectedHeader]

			Expect(headerValue).To(Equal(expectedHeaderValue))
		})

		It("should create a Predicate with Query parameters", func() {
			containsPredicate := actualPredicateAsMap["contains"]

			Expect(containsPredicate).To(HaveKey("query"))
		})

		It("should create a Predicate with the correct Query parameter", func() {
			containsPredicate := actualPredicateAsMap["contains"].(map[string]interface{})

			queryParametersMap := containsPredicate["query"].(map[string]interface{})

			Expect(queryParametersMap).To(HaveKeyWithValue(expectedQueryParam, expectedQueryParamValue))
		})

		It("should create a Predicate with the correct Body", func() {
			containsPredicate := actualPredicateAsMap["contains"]

			Expect(containsPredicate).To(HaveKeyWithValue("body", expectedBody))
		})

	})

})
