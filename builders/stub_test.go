package builders_test

import (
	. "github.com/durmaze/gobank/builders"
	"github.com/durmaze/gobank/predicates"
	"github.com/durmaze/gobank/responses"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"sync"
)

var _ = Describe("Stub Builder Tests", func() {

	Describe("When building a Stub with single Response", func() {
		var (
			stub Stub
			once sync.Once
		)

		BeforeEach(func() {
			once.Do(func() {

				expectedResponse := responses.Is().StatusCode(200).Body("{ \"greeting\": \"Hello GoBank\" }").Build()
				stub = NewStubBuilder().Responses(expectedResponse).Build()

			})
		})

		It("should create a Stub that has one Predicate", func() {
			Expect(stub.Responses).To(HaveLen(1))
		})

		It("should create a Stub that has a Predicate with type \"Equals\"", func() {
			Expect(stub.Responses[0].Type()).To(Equal(responses.Is().Build().Type()))
		})
	})

	Describe("When building a Stub with single Response and a single Equals predicate", func() {
		var (
			stub Stub

			once sync.Once
		)

		BeforeEach(func() {
			once.Do(func() {

				expectedResponse := responses.Is().StatusCode(200).Body("{ \"greeting\": \"Hello GoBank\" }").Build()
				expectedPredicate := predicates.Equals().Path("/test-path").Build()

				stub = NewStubBuilder().Responses(expectedResponse).Predicates(expectedPredicate).Build()
			})
		})

		It("should create a Stub that has one Predicate", func() {
			Expect(stub.Responses).To(HaveLen(1))
		})

		It("should create a Stub that has a Predicate with type \"Equals\"", func() {
			Expect(stub.Responses[0].Type()).To(Equal(responses.Is().Build().Type()))
		})

		It("should create a Stub that has one Predicate", func() {
			Expect(stub.Predicates).To(HaveLen(1))
		})

		It("should create a Stub that has a Predicate with type \"Equals\"", func() {
			Expect(stub.Predicates[0].Type()).To(Equal(predicates.Equals().Build().Type()))
		})
	})

	Describe("When building a Stub with single Response and multiple different predicates", func() {
		var (
			stub Stub
			once sync.Once
		)

		BeforeEach(func() {
			once.Do(func() {

				expectedResponse := responses.Is().StatusCode(200).Body("{ \"greeting\": \"Hello GoBank\" }").Build()

				expectedPredicate1 := predicates.Equals().Path("/test-path").Build()
				expectedPredicate2 := predicates.Contains().Method("POST").Build()

				stub = NewStubBuilder().Responses(expectedResponse).Predicates(expectedPredicate1, expectedPredicate2).Build()

			})
		})

		It("should create a Stub that has one Predicate", func() {
			Expect(stub.Responses).To(HaveLen(1))
		})

		It("should create a Stub that has a Predicate with type \"Equals\"", func() {
			Expect(stub.Responses[0].Type()).To(Equal(responses.Is().Build().Type()))
		})

		It("should create a Stub that has two Predicates", func() {
			Expect(stub.Predicates).To(HaveLen(2))
		})

		It("should create a Stub that has a Predicate with type \"Equals\"", func() {
			Expect(stub.Predicates[0].Type()).To(Equal(predicates.Equals().Build().Type()))
		})

		It("should create a Stub that has a Predicate with type \"Contains\"", func() {
			Expect(stub.Predicates[1].Type()).To(Equal(predicates.Contains().Build().Type()))
		})

	})

})
