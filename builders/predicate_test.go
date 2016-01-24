package builders_test

import (
	"sync"

	. "github.com/durmaze/gobank/builders"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Predicate Builder Tests", func() {

	Describe("When building a Predicate of type \"Equals\"", func() {

		var (
			actualPredicate Predicate

			expectedPath            = "/testing/path"
			expectedMethod          = "GET"
			expectedHeader          = "Content-Type"
			expectedHeaderValue     = "application/json"
			expectedBody            = "{ \"greeting\": \"Hello GoBank\" }"
			expectedQueryParam      = "qp"
			expectedQueryParamValue = "testValue"

			once sync.Once
		)

		BeforeEach(func() {
			once.Do(func() {
				actualPredicate = NewPredicateBuilder().
					EqualsPredicate().
					Path(expectedPath).
					Method(expectedMethod).
					Header(expectedHeader, expectedHeaderValue).
					Query(expectedQueryParam, expectedQueryParamValue).
					Body(expectedBody).
					Build()
			})
		})

		It("should create a Predicate with the correct Path", func() {
			Expect(actualPredicate.Equals.Path).To(Equal(expectedPath))
		})

		It("should create a Predicate with the correct Method", func() {
			Expect(actualPredicate.Equals.Method).To(Equal(expectedMethod))
		})

		It("should create a Predicate with the correct Header", func() {
			Expect(actualPredicate.Equals.Headers[expectedHeader]).To(Equal(expectedHeaderValue))
		})

		It("should create a Predicate with the correct Query parameter", func() {
			Expect(actualPredicate.Equals.QueryParams[expectedQueryParam]).To(Equal(expectedQueryParamValue))
		})

		It("should create a Predicate with the correct Body", func() {
			Expect(actualPredicate.Equals.Body).To(Equal(expectedBody))
		})

	})

})
