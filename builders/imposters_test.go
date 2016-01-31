package builders_test

import (
	. "github.com/durmaze/gobank/builders"
	"github.com/durmaze/gobank/predicates"
	"github.com/durmaze/gobank/responses"

	"sync"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Imposter Builder Tests", func() {

	Describe("When building an imposter with protocol, port and multiple stubs", func() {
		var (
			imposter Imposter
			once     sync.Once
		)

		BeforeEach(func() {
			once.Do(func() {
				response := responses.Is().StatusCode(200).Body("{}").Build()

				equals := predicates.Equals().Path("/test-path").Build()
				contains := predicates.Contains().Header("Content-Type", "application/json").Build()

				stub := NewStubBuilder().Responses(response).Predicates(equals, contains).Build()

				imposter = NewImposterBuilder().Protocol("http").Port(4546).Stubs(stub).Build()
			})
		})

		It("should create the Imposter with the specified protocol", func() {
			Expect(imposter.Protocol).To(Equal("http"))
		})

		It("should create the Imposter on the specified port", func() {
			Expect(imposter.Port).To(Equal(4546))
		})
	})

})
