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
			actualResponse Response
			expectedResponse Response
			once sync.Once
		)

		BeforeEach(func() {
			once.Do(func(){
				is := Is{
					StatusCode: 200,
				}

				expectedResponse = Response{
					Is: is,
				}

				stub := NewStubBuilder().AddResponse(expectedResponse).Build()

				actualResponse = stub.Responses[0]
			})
	  })

		It("should create a Stub that returns a Response with the correct StatusCode", func() {
			Expect(actualResponse.Is.StatusCode).To(Equal(expectedResponse.Is.StatusCode))
		})
	})

})
