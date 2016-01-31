package predicates_test

import (
	"encoding/json"
	"log"
	"sync"

	"github.com/durmaze/gobank/predicates"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Or Predicate Builder Tests", func() {
	Describe("When building a Predicate of type \"Or\"", func() {

		var (
			actualPredicateAsMap map[string]interface{}
			predicate1           = predicates.NewEqualsBuilder().Build()
			predicate2           = predicates.NewContainsBuilder().Build()
			predicate3           = predicates.NewEqualsBuilder().Build()

			once sync.Once
		)
		BeforeEach(func() {
			once.Do(func() {
				actualPredicate := predicates.NewOrBuilder().AddPredicate(predicate1).AddPredicate(predicate2).AddPredicate(predicate3).Build()
				log.Println("actualPredicate", actualPredicate)

				jsonBytes, err := json.Marshal(actualPredicate)
				log.Println("jsonBytes", string(jsonBytes))
				log.Println("err", err)

				actualPredicateAsMap = map[string]interface{}{}
				json.Unmarshal(jsonBytes, &actualPredicateAsMap)

				log.Println("predicateMap", actualPredicateAsMap)
				log.Println("orPredicate", actualPredicateAsMap["or"])
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
