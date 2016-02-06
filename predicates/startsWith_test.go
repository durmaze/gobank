package predicates_test

import (
	"encoding/json"
	"sync"

	"github.com/durmaze/gobank/predicates"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("StartsWith Predicate Builder Tests", func() {

	Describe("When building a Predicate of type \"StartsWith\"", func() {

		var (
			actualPredicateAsMap map[string]interface{}

			expectedPath = "/testing/"

			once sync.Once
		)

		BeforeEach(func() {
			once.Do(func() {
				actualPredicate := predicates.StartsWith().Path(expectedPath).Build()

				jsonBytes, _ := json.Marshal(actualPredicate)
				actualPredicateAsMap = map[string]interface{}{}
				json.Unmarshal(jsonBytes, &actualPredicateAsMap)
			})
		})

		It("should create a \"StartsWith\" Predicate", func() {
			Expect(actualPredicateAsMap).To(HaveKey("startsWith"))
		})

		It("should create a Predicate with the correct Path", func() {
			startsWithPredicate := actualPredicateAsMap["startsWith"]

			Expect(startsWithPredicate).To(HaveKeyWithValue("path", expectedPath))
		})

	})

})
