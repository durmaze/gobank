package builders_test

import (
	. "github.com/durmaze/gobank/builders"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"sync"
)

var _ = Describe("Stub Builder Tests", func() {

	Describe("When building a Stub with single Response", func() {
		var (
			actualResponse   Response
			expectedResponse Response
			once             sync.Once
		)

		BeforeEach(func() {
			once.Do(func() {

				expectedResponse = Response{
					Is: Is{
						StatusCode: 200,
						Body:       "{ \"greeting\": \"Hello GoBank\" }",
					},
				}

				stub := NewStubBuilder().AddResponse(expectedResponse).Build()

				actualResponse = stub.Responses[0]
			})
		})

		It("should create a Stub that returns a Response with the correct StatusCode", func() {
			Expect(actualResponse.Is.StatusCode).To(Equal(expectedResponse.Is.StatusCode))
		})

		It("should create a Stub that returns a Response with the correct Body", func() {
			Expect(actualResponse.Is.Body).To(Equal(expectedResponse.Is.Body))
		})
	})

	Describe("When building a Stub with single Response and a single Equals predicate", func() {
		var (
			actualResponse   Response
			expectedResponse Response

			actualPredicate   Predicate
			expectedPredicate Predicate

			once sync.Once
		)

		BeforeEach(func() {
			once.Do(func() {

				expectedResponse = Response{
					Is: Is{
						StatusCode: 200,
						Body:       "{ \"greeting\": \"Hello GoBank\" }",
					},
				}

				expectedPredicate = Predicate{
					Equals: Equals{
						Path: "/test-path",
					},
				}

				stub := NewStubBuilder().AddResponse(expectedResponse).AddPredicate(expectedPredicate).Build()

				actualResponse = stub.Responses[0]
				actualPredicate = stub.Predicates[0]
			})
		})

		It("should create a Stub that returns a Response with the correct StatusCode", func() {
			Expect(actualResponse.Is.StatusCode).To(Equal(expectedResponse.Is.StatusCode))
		})

		It("should create a Stub that returns a Response with the correct Body", func() {
			Expect(actualResponse.Is.Body).To(Equal(expectedResponse.Is.Body))
		})

		It("should create a Stub that has a Predicate with type \"Equals\"", func() {
			Expect(actualPredicate.Equals).NotTo(BeNil())
		})

		It("should create a Stub that has a Predicate with correct Path condition", func() {
			Expect(actualPredicate.Equals.Path).To(Equal(expectedPredicate.Equals.Path))
		})

	})

	Describe("When building a Stub with single Response and multiple Equals predicates", func() {
		var (
			actualResponse   Response
			expectedResponse Response

			actualPredicate1    Predicate
			actualPredicate2    Predicate			
			expectedPredicate1 Predicate
			expectedPredicate2 Predicate

			once sync.Once
		)

		BeforeEach(func() {
			once.Do(func() {

				expectedResponse = Response{
					Is: Is{
						StatusCode: 200,
						Body:       "{ \"greeting\": \"Hello GoBank\" }",
					},
				}

				expectedPredicate1 = Predicate{
					Equals: Equals{
						Path: "/test-path",
					},
				}


				expectedPredicate2 = Predicate{
					Equals: Equals{
						Method: "POST",
					},
				}

				stub := NewStubBuilder().AddResponse(expectedResponse).AddPredicate(expectedPredicate1).AddPredicate(expectedPredicate2).Build()

				actualResponse = stub.Responses[0]
				actualPredicate1 = stub.Predicates[0]
				actualPredicate2 = stub.Predicates[1]
			})
		})

		It("should create a Stub that returns a Response with the correct StatusCode", func() {
			Expect(actualResponse.Is.StatusCode).To(Equal(expectedResponse.Is.StatusCode))
		})

		It("should create a Stub that returns a Response with the correct Body", func() {
			Expect(actualResponse.Is.Body).To(Equal(expectedResponse.Is.Body))
		})

		It("should create a Stub that has 2 Predicates of type \"Equals\"", func() {
			Expect(actualPredicate1.Equals).NotTo(BeNil())
			Expect(actualPredicate2.Equals).NotTo(BeNil())			
		})

		It("should create a Stub that has a Predicate with correct Path condition", func() {
			Expect(actualPredicate1.Equals.Path).To(Equal(expectedPredicate1.Equals.Path))
		})

		It("should create a Stub that has a Predicate with correct Method condition", func() {
			Expect(actualPredicate2.Equals.Method).To(Equal(expectedPredicate2.Equals.Method))
		})

	})

})
