package predicates_test

import (
	"encoding/json"
	"log"
	"sync"

	"github.com/durmaze/gobank/predicates"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Exists Predicate Builder Tests", func() {

	Describe("When building a Predicate of type \"Exists\"", func() {

		var (
			actualPredicateAsMap map[string]interface{}

			expectedPath             = true
			expectedHeader           = "Accept"
			expectedHeaderValue      = true
			expectedBody             = false
			expectedQueryParam1      = "q"
			expectedQueryParamValue1 = true
			expectedQueryParam2      = "search"
			expectedQueryParamValue2 = false

			once sync.Once
		)

		BeforeEach(func() {
			once.Do(func() {
				actualPredicate := predicates.Exists().
					Path(expectedPath).
					Header(expectedHeader, expectedHeaderValue).
					Query(expectedQueryParam1, expectedQueryParamValue1).
					Query(expectedQueryParam2, expectedQueryParamValue2).
					Body(expectedBody).
					Build()
				log.Println(actualPredicate, "******************")
				jsonBytes, _ := json.Marshal(actualPredicate)

				actualPredicateAsMap = map[string]interface{}{}
				json.Unmarshal(jsonBytes, &actualPredicateAsMap)
				log.Println(actualPredicateAsMap, "***************")
			})
		})

		It("should create a \"Exists\" Predicate", func() {
			Expect(actualPredicateAsMap).To(HaveKey("exists"))
		})

		It("should create a Predicate with the correct Path", func() {
			existsPredicate := actualPredicateAsMap["exists"]

			Expect(existsPredicate).To(HaveKeyWithValue("path", expectedPath))
		})

		It("should create a Predicate without a Method", func() {
			existsPredicate := actualPredicateAsMap["exists"]

			Expect(existsPredicate).NotTo(HaveKey("method"))
		})

		It("should create a Predicate with Headers", func() {
			existsPredicate := actualPredicateAsMap["exists"]

			Expect(existsPredicate).To(HaveKey("headers"))
		})

		It("should create a Predicate with the correct Header", func() {
			existsPredicate := actualPredicateAsMap["exists"].(map[string]interface{})

			headersMap := existsPredicate["headers"].(map[string]interface{})
			headerValue := headersMap[expectedHeader]

			Expect(headerValue).To(Equal(expectedHeaderValue))
		})

		It("should create a Predicate with Query parameters", func() {
			existsPredicate := actualPredicateAsMap["exists"]

			Expect(existsPredicate).To(HaveKey("query"))
		})

		It("should create a Predicate with the correct Query parameter", func() {
			existsPredicate := actualPredicateAsMap["exists"].(map[string]interface{})

			queryParametersMap := existsPredicate["query"].(map[string]interface{})

			Expect(queryParametersMap).To(HaveKeyWithValue(expectedQueryParam1, expectedQueryParamValue1))
			Expect(queryParametersMap).To(HaveKeyWithValue(expectedQueryParam2, expectedQueryParamValue2))
		})

		It("should create a Predicate with the correct Body", func() {
			existsPredicate := actualPredicateAsMap["exists"]

			Expect(existsPredicate).To(HaveKeyWithValue("body", expectedBody))
		})

	})

})
