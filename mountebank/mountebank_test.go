package mountebank_test

import (
	"net/http"
	"strconv"

	"github.com/durmaze/gobank/builders"
	"github.com/durmaze/gobank/mountebank"
	"github.com/durmaze/gobank/predicates"
	"github.com/durmaze/gobank/responses"
	"github.com/parnurzeal/gorequest"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Mountebank Client", func() {

	Describe("When createImposter request is sent to Mountebank", func() {

		var (
			protocol = "http"
			port     = 4546
		)

		BeforeEach(func() {
			okResponse := responses.Is().StatusCode(200).Body("{ \"greeting\": \"Hello GoBank\" }").Build()

			equals := predicates.Equals().Path("/test-path").Build()
			contains := predicates.Contains().Header("Accept", "application/json").Build()
			or := predicates.Or().Predicates(equals, contains).Build()

			stub := builders.Stub().Responses(okResponse).Predicates(or).Build()

			imposter := builders.NewImposterBuilder().Protocol(protocol).Port(port).Stubs(stub).Build()

			client := mountebank.NewClient(MountebankUri)
			client.CreateImposter(imposter)
		})

		It("should have the Imposter installed on Mountebank", func() {
			imposterUri := MountebankUri + "/imposters/" + strconv.Itoa(port)
			resp, _, _ := gorequest.New().Get(imposterUri).End()

			Expect(resp.StatusCode).To(Equal(http.StatusOK))
		})
	})

	Describe("When deleteImposter request is sent to Mountebank", func() {

		var (
			protocol = "http"
			port     = 5000
		)

		BeforeEach(func() {
			imposter := builders.NewImposterBuilder().Protocol(protocol).Port(port).Build()
			client := mountebank.NewClient(MountebankUri)
			client.CreateImposter(imposter)

			client.DeleteImposter(imposter)
		})

		It("should delete the installed Imposter at Mountebank", func() {
			imposterUri := MountebankUri + "/imposters/" + strconv.Itoa(port)
			resp, _, _ := gorequest.New().Get(imposterUri).End()

			Expect(resp.StatusCode).To(Equal(http.StatusNotFound))
		})

	})

	Describe("When deleteAllImposter request is sent to Mountebank", func() {

		BeforeEach(func() {
			client := mountebank.NewClient(MountebankUri)
			client.DeleteAllImposters()
		})

		It("should delete all the installed Imposters at Mountebank", func() {
			impostersUri := MountebankUri + "/imposters"
			resp, _, _ := gorequest.New().Get(impostersUri).End()

			Expect(resp.StatusCode).To(Equal(http.StatusOK))
		})

	})

})
