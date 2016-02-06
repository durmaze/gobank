package predicates_test

import (
	"encoding/json"
	"sync"

	"github.com/durmaze/gobank/predicates"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Not Predicate Builder Tests", func() {
	Describe("When building a Predicate of type \"Not\"", func() {

		var (
			actualPredicateAsMap map[string]interface{}
			predicate            = predicates.Equals().Build()

			once sync.Once
		)
		BeforeEach(func() {
			once.Do(func() {
				actualPredicate := predicates.Not().Predicate(predicate).Build()

				jsonBytes, _ := json.Marshal(actualPredicate)

				actualPredicateAsMap = map[string]interface{}{}
				json.Unmarshal(jsonBytes, &actualPredicateAsMap)

			})
		})

		It("should create an \"Not\" Predicate", func() {
			Expect(actualPredicateAsMap).To(HaveKey("not"))
		})

		It("should have the child predicate", func() {
			childPredicate := actualPredicateAsMap["not"].(interface{})

			Expect(childPredicate).To(HaveKey("equals"))
		})

	})
})
