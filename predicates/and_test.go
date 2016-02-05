package predicates_test

import (
	"encoding/json"
	"sync"

	"github.com/durmaze/gobank/predicates"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("And Predicate Builder Tests", func() {
	Describe("When building a Predicate of type \"And\"", func() {

		var (
			actualPredicateAsMap map[string]interface{}
			predicate1           = predicates.Equals().Build()
			predicate2           = predicates.Contains().Build()
			predicate3           = predicates.Equals().Build()

			once sync.Once
		)
		BeforeEach(func() {
			once.Do(func() {
				actualPredicate := predicates.And().Predicates(predicate1, predicate2, predicate3).Build()

				jsonBytes, _ := json.Marshal(actualPredicate)

				actualPredicateAsMap = map[string]interface{}{}
				json.Unmarshal(jsonBytes, &actualPredicateAsMap)

			})
		})

		It("should create an \"And\" Predicate", func() {
			Expect(actualPredicateAsMap).To(HaveKey("and"))
		})

		It("should have list of predicates", func() {
			andPredicate := actualPredicateAsMap["and"]

			Expect(andPredicate).To(HaveLen(3))
		})

		It("should have the first predicate", func() {
			childPredicates := actualPredicateAsMap["and"].([]interface{})

			Expect(childPredicates[0]).To(HaveKey("equals"))
		})

		It("should have the second predicate", func() {
			childPredicates := actualPredicateAsMap["and"].([]interface{})

			Expect(childPredicates[1]).To(HaveKey("contains"))
		})

		It("should have the third predicate", func() {
			childPredicates := actualPredicateAsMap["and"].([]interface{})

			Expect(childPredicates[2]).To(HaveKey("equals"))
		})

	})
})
