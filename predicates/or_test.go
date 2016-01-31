package predicates_test

import (
	"encoding/json"
	"sync"

	"github.com/durmaze/gobank/predicates"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Or Predicate Builder Tests", func() {
	Describe("When building a Predicate of type \"Or\"", func() {

		var (
			actualPredicateAsMap map[string]interface{}
			predicate1           = predicates.Equals().Build()
			predicate2           = predicates.Contains().Build()
			predicate3           = predicates.Equals().Build()

			once sync.Once
		)
		BeforeEach(func() {
			once.Do(func() {
				actualPredicate := predicates.Or().Predicates(predicate1, predicate2, predicate3).Build()

				jsonBytes, _ := json.Marshal(actualPredicate)

				actualPredicateAsMap = map[string]interface{}{}
				json.Unmarshal(jsonBytes, &actualPredicateAsMap)

			})
		})

		It("should create an \"Or\" Predicate", func() {
			Expect(actualPredicateAsMap).To(HaveKey("or"))
		})

		It("should have list of predicates", func() {
			orPredicate := actualPredicateAsMap["or"]

			Expect(orPredicate).To(HaveLen(3))
		})

		It("should have the first predicate", func() {
			childPredicates := actualPredicateAsMap["or"].([]interface{})

			Expect(childPredicates[0]).To(HaveKey("equals"))
		})

		It("should have the second predicate", func() {
			childPredicates := actualPredicateAsMap["or"].([]interface{})

			Expect(childPredicates[1]).To(HaveKey("contains"))
		})

		It("should have the third predicate", func() {
			childPredicates := actualPredicateAsMap["or"].([]interface{})

			Expect(childPredicates[2]).To(HaveKey("equals"))
		})

	})
})
