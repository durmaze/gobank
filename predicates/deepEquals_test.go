package predicates_test

import (
	"encoding/json"
	"sync"

	"github.com/durmaze/gobank/predicates"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("DeepEquals Predicate Builder Tests", func() {

	Describe("When building a Predicate of type \"DeepEquals\"", func() {

		var (
			actualPredicateAsMap map[string]interface{}

			expectedPath            = "/testing/path"
			expectedMethod          = "GET"
			expectedQueryParam      = "qp"
			expectedQueryParamValue = "testValue"

			once sync.Once
		)

		BeforeEach(func() {
			once.Do(func() {
				actualPredicate := predicates.DeepEquals().
					Path(expectedPath).
					Method(expectedMethod).
					Query(expectedQueryParam, expectedQueryParamValue).
					Build()

				jsonBytes, _ := json.Marshal(actualPredicate)
				actualPredicateAsMap = map[string]interface{}{}
				json.Unmarshal(jsonBytes, &actualPredicateAsMap)
			})
		})

		It("should create a \"DeepEquals\" Predicate", func() {
			Expect(actualPredicateAsMap).To(HaveKey("deepEquals"))
		})

		It("should create a Predicate with the correct Path", func() {
			deepEqualsPredicate := actualPredicateAsMap["deepEquals"]

			Expect(deepEqualsPredicate).To(HaveKeyWithValue("path", expectedPath))
		})

		It("should create a Predicate with the correct Method", func() {
			deepEqualsPredicate := actualPredicateAsMap["deepEquals"]

			Expect(deepEqualsPredicate).To(HaveKeyWithValue("method", expectedMethod))
		})

		It("should create a Predicate with Query parameters", func() {
			deepEqualsPredicate := actualPredicateAsMap["deepEquals"]

			Expect(deepEqualsPredicate).To(HaveKey("query"))
		})

		It("should create a Predicate with the correct Query parameter", func() {
			deepEqualsPredicate := actualPredicateAsMap["deepEquals"].(map[string]interface{})

			queryParametersMap := deepEqualsPredicate["query"].(map[string]interface{})

			Expect(queryParametersMap).To(HaveKeyWithValue(expectedQueryParam, expectedQueryParamValue))
		})

	})

})
