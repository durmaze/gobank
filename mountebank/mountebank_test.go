package mountebank_test

import (
	"net/http"

	"github.com/durmaze/gobank/builders"
	"github.com/durmaze/gobank/mountebank"
	"github.com/parnurzeal/gorequest"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("When create Imposter request is sent to Mountebank", func() {

	BeforeEach(func() {
		imposterBuilder := builders.NewImposterBuilder()
		imposter := imposterBuilder.Protocol("http").Port(4546).Build()

		mountebank.CreateImposter(imposter)
	})

	It("should have the Imposter installed on Mountebank", func() {
		resp, _, _ := gorequest.New().Get("http://localhost:2525/imposters/4546").End()

		Expect(resp.StatusCode).To(Equal(http.StatusOK))
	})

})
