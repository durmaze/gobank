package mountebank_test

import (
	"net/http"

	"github.com/durmaze/gobank/builders"
	"github.com/durmaze/gobank/mountebank"
	"github.com/durmaze/gobank/predicates"
	"github.com/durmaze/gobank/responses"
	"github.com/parnurzeal/gorequest"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("When create Imposter request is sent to Mountebank", func() {

	BeforeEach(func() {

		expectedResponse := responses.Is().StatusCode(200).Body("{ \"greeting\": \"Hello GoBank\" }").Build()

		expectedPredicate1 := predicates.Equals().Path("/test-path").Build()
		expectedPredicate2 := predicates.Contains().Header("Content-Type", "application/json").Build()

		orPredicate := predicates.Or().Predicates(expectedPredicate1, expectedPredicate2).Build()
		stub1 := builders.Stub().Responses(expectedResponse).Predicates(orPredicate).Build()

		imposterBuilder := builders.NewImposterBuilder()
		imposter := imposterBuilder.Protocol("http").Port(4546).Stubs(stub1).Build()

		mountebank.CreateImposter(imposter)
	})

	It("should have the Imposter installed on Mountebank", func() {
		resp, _, _ := gorequest.New().Get("http://localhost:2525/imposters/4546").End()

		Expect(resp.StatusCode).To(Equal(http.StatusOK))
	})

})
