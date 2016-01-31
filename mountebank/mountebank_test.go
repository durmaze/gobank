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
			port = 4546
		)

		BeforeEach(func() {
				expectedResponse := responses.Is().StatusCode(200).Body("{ \"greeting\": \"Hello GoBank\" }").Build()

				expectedPredicate1 := predicates.Equals().Path("/test-path").Build()
				expectedPredicate2 := predicates.Contains().Header("Content-Type", "application/json").Build()

				orPredicate := predicates.Or().Predicates(expectedPredicate1, expectedPredicate2).Build()
				stub1 := builders.Stub().Responses(expectedResponse).Predicates(orPredicate).Build()

				imposter := builders.NewImposterBuilder().Protocol(protocol).Port(port).Stubs(stub1).Build()

				mountebank.CreateImposter(imposter)
			})

			It("should have the Imposter installed on Mountebank", func() {
				imposterUri := MountebankBaseUri + "/imposters/" + strconv.Itoa(port)
				resp, _, _ := gorequest.New().Get(imposterUri).End()

				Expect(resp.StatusCode).To(Equal(http.StatusOK))
			})
	})

	Describe("When deleteImposter request is sent to Mountebank", func() {
		
		var (
			protocol = "http"
			port = 5000
		)

		BeforeEach(func() {
			imposter := builders.NewImposterBuilder().Protocol(protocol).Port(port).Build()
			mountebank.CreateImposter(imposter)

			mountebank.DeleteImposter(imposter)
		})

		It("should delete the installed Imposter at Mountebank", func() {
			imposterUri := MountebankBaseUri + "/imposters/" + strconv.Itoa(port)
			resp, _, _ := gorequest.New().Get(imposterUri).End()

			Expect(resp.StatusCode).To(Equal(http.StatusNotFound))
		})

	})

	Describe("When deleteAllImposter request is sent to Mountebank", func() {

		BeforeEach(func() {
			mountebank.DeleteAllImposters()
		})

		It("should delete all the installed Imposters at Mountebank", func() {
			impostersUri := MountebankBaseUri + "/imposters"
			resp, _, _ := gorequest.New().Get(impostersUri).End()

			Expect(resp.StatusCode).To(Equal(http.StatusOK))
		})

	})

})
