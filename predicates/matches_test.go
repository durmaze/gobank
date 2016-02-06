package predicates_test

import (
	"encoding/json"
	"sync"

	"github.com/durmaze/gobank/predicates"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Matches Predicate Builder Tests", func() {

	Describe("When building a Predicate of type \"Matches\"", func() {

		var (
			actualPredicateAsMap map[string]interface{}

			expectedPath = "^/testing/"

			once sync.Once
		)

		BeforeEach(func() {
			once.Do(func() {
				actualPredicate := predicates.Matches().Path(expectedPath).Build()

				jsonBytes, _ := json.Marshal(actualPredicate)
				actualPredicateAsMap = map[string]interface{}{}
				json.Unmarshal(jsonBytes, &actualPredicateAsMap)
			})
		})

		It("should create a \"Matches\" Predicate", func() {
			Expect(actualPredicateAsMap).To(HaveKey("matches"))
		})

		It("should create a Predicate with the correct Path", func() {
			matchesPredicate := actualPredicateAsMap["matches"]

			Expect(matchesPredicate).To(HaveKeyWithValue("path", expectedPath))
		})

	})

})
