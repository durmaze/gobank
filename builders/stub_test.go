package builders_test

import (
	"fmt"
	"log"

	. "github.com/durmaze/gobank/builders"
	"github.com/durmaze/gobank/predicates"
	. "github.com/durmaze/gobank/responses"

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

			stub Stub

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
				log.Println("expectedResponse", expectedResponse)
				expectedPredicate := predicates.NewEqualsBuilder().Path("/test-path").Build()
				log.Println("expectedPredicate", expectedPredicate)

				stub = NewStubBuilder().AddResponse(expectedResponse).AddPredicate(expectedPredicate).Build()

				actualResponse = stub.Responses[0]
			})
		})

		It("should create a Stub that returns a Response with the correct StatusCode", func() {
			Expect(actualResponse.Is.StatusCode).To(Equal(expectedResponse.Is.StatusCode))
		})

		It("should create a Stub that returns a Response with the correct Body", func() {
			Expect(actualResponse.Is.Body).To(Equal(expectedResponse.Is.Body))
		})

		It("should create a Stub that has one Predicate", func() {
			Expect(stub.Predicates).To(HaveLen(1))
		})

		It("should create a Stub that has a Predicate with type \"Equals\"", func() {
			Expect(TypeOf(stub.Predicates[0])).To(Equal(TypeOf(predicates.Equals{})))
		})
	})

	Describe("When building a Stub with single Response and multiple different predicates", func() {
		var (
			actualResponse   Response
			expectedResponse Response

			stub Stub
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

				expectedPredicate1 := predicates.NewEqualsBuilder().Path("/test-path").Build()

				expectedPredicate2 := predicates.NewContainsBuilder().Method("POST").Build()

				stub = NewStubBuilder().AddResponse(expectedResponse).AddPredicate(expectedPredicate1).AddPredicate(expectedPredicate2).Build()

				actualResponse = stub.Responses[0]
			})
		})

		It("should create a Stub that returns a Response with the correct StatusCode", func() {
			Expect(actualResponse.Is.StatusCode).To(Equal(expectedResponse.Is.StatusCode))
		})

		It("should create a Stub that returns a Response with the correct Body", func() {
			Expect(actualResponse.Is.Body).To(Equal(expectedResponse.Is.Body))
		})

		It("should create a Stub that has two Predicates", func() {
			Expect(stub.Predicates).To(HaveLen(2))
		})

		It("should create a Stub that has a Predicate with type \"Equals\"", func() {
			Expect(TypeOf(stub.Predicates[0])).To(Equal(TypeOf(predicates.Equals{})))
		})

		It("should create a Stub that has a Predicate with type \"Contains\"", func() {
			Expect(TypeOf(stub.Predicates[1])).To(Equal(TypeOf(predicates.Contains{})))
		})

	})

})

func TypeOf(instance interface{}) string {
	return fmt.Sprintf("%T", instance)
}
