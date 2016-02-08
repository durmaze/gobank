package predicates_test

import (
	"encoding/json"
	"sync"

	"github.com/durmaze/gobank/predicates"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Inject Predicate Builder Tests", func() {

	Describe("When building a Predicate of type \"Inject\"", func() {

		var (
			actualPredicateAsMap map[string]interface{}

			expectedFunction = "function (request) { return new Buffer(request.data, 'base64')[2] <= 100; }"

			once sync.Once
		)

		BeforeEach(func() {
			once.Do(func() {
				actualPredicate := predicates.Inject().Function(expectedFunction).Build()

				jsonBytes, _ := json.Marshal(actualPredicate)
				actualPredicateAsMap = map[string]interface{}{}
				json.Unmarshal(jsonBytes, &actualPredicateAsMap)
			})
		})

		It("should create a \"Inject\" Predicate with the correct function", func() {
			Expect(actualPredicateAsMap).To(HaveKeyWithValue("inject", expectedFunction))
		})

	})

})
