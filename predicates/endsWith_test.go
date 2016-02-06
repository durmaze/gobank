package predicates_test

import (
	"encoding/json"
	"sync"

	"github.com/durmaze/gobank/predicates"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("EndsWith Predicate Builder Tests", func() {

	Describe("When building a Predicate of type \"EndsWith\"", func() {

		var (
			actualPredicateAsMap map[string]interface{}

			expectedPath = ".jpeg"

			once sync.Once
		)

		BeforeEach(func() {
			once.Do(func() {
				actualPredicate := predicates.EndsWith().Path(expectedPath).Build()

				jsonBytes, _ := json.Marshal(actualPredicate)
				actualPredicateAsMap = map[string]interface{}{}
				json.Unmarshal(jsonBytes, &actualPredicateAsMap)
			})
		})

		It("should create a \"EndsWith\" Predicate", func() {
			Expect(actualPredicateAsMap).To(HaveKey("endsWith"))
		})

		It("should create a Predicate with the correct Path", func() {
			endsWithPredicate := actualPredicateAsMap["endsWith"]

			Expect(endsWithPredicate).To(HaveKeyWithValue("path", expectedPath))
		})

	})

})
