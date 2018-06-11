package gobank_test

import (
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/durmaze/gobank"
	"github.com/durmaze/gobank/predicates"
	"github.com/durmaze/gobank/responses"
	"github.com/parnurzeal/gorequest"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Mountebank Client", func() {

	Describe("When createImposter request is sent to Mountebank", func() {

		var (
			protocol         = "http"
			port             = 4546
			createdImposter  map[string]interface{}
			expectedImposter gobank.ImposterElement
			err              error

			once sync.Once
		)

		BeforeEach(func() {
			once.Do(func() {
				okResponse := responses.Is().StatusCode(200).Body("{ \"greeting\": \"Hello GoBank\" }").Build()

				equals := predicates.Equals().Path("/test-path").Build()
				contains := predicates.Contains().Header("Accept", "application/json").Build()
				exists := predicates.Exists().Method(true).Query("q", false).Body(false).Build()
				or := predicates.Or().Predicates(equals, contains, exists).Build()

				stub := gobank.Stub().Responses(okResponse).Predicates(or).Build()

				expectedImposter = gobank.NewImposterBuilder().Protocol(protocol).Port(port).Name("Greeting Imposter").Stubs(stub).Build()

				client := gobank.NewClient(MountebankURI)
				createdImposter, err = client.CreateImposter(expectedImposter)
				log.Println("ActualImposter: ", createdImposter)
			})
		})

		It("should have the Imposter installed on Mountebank", func() {
			imposterURI := MountebankURI + "/imposters/" + strconv.Itoa(port)
			resp, body, _ := gorequest.New().Get(imposterURI).End()

			log.Println("Imposter from Mountebank. Body: ", body)
			Expect(resp.StatusCode).To(Equal(http.StatusOK))
		})

		It("should return the correct protocol", func() {
			Expect(createdImposter["protocol"]).To(Equal(protocol))
		})

		It("should return the correct port", func() {
			Expect(createdImposter["port"]).To(Equal(float64(port)))
		})

		It("should return the correct name", func() {
			Expect(createdImposter["name"]).To(Equal("Greeting Imposter"))
		})

		It("should return one stub", func() {
			Expect(createdImposter["stubs"]).To(HaveLen(1))
		})

		It("should not return an error", func() {
			Expect(err).To(BeNil())
		})
	})

	Describe("When deleteImposter request is sent to Mountebank", func() {

		var (
			protocol        = "http"
			port            = 5000
			deletedImposter map[string]interface{}
			err             error

			once sync.Once
		)

		BeforeEach(func() {
			once.Do(func() {
				imposter := gobank.NewImposterBuilder().Protocol(protocol).Port(port).Build()
				client := gobank.NewClient(MountebankURI)
				client.CreateImposter(imposter)

				deletedImposter, err = client.DeleteImposter(port)
			})
		})

		It("should delete the installed Imposter at Mountebank", func() {
			imposterURI := MountebankURI + "/imposters/" + strconv.Itoa(port)
			resp, _, _ := gorequest.New().Get(imposterURI).End()

			Expect(resp.StatusCode).To(Equal(http.StatusNotFound))
		})

		It("should return the correct protocol", func() {
			Expect(deletedImposter["protocol"]).To(Equal(protocol))
		})

		It("should return the correct port", func() {
			Expect(deletedImposter["port"]).To(Equal(float64(port)))
		})

		It("should not return an error", func() {
			Expect(err).To(BeNil())
		})
	})

	Describe("When deleteAllImposter request is sent to Mountebank", func() {
		var (
			protocol = "http"
			err      error

			once sync.Once
		)

		BeforeEach(func() {
			once.Do(func() {
				imposter1 := gobank.NewImposterBuilder().Protocol(protocol).Build()
				imposter2 := gobank.NewImposterBuilder().Protocol(protocol).Build()

				client := gobank.NewClient(MountebankURI)
				client.CreateImposter(imposter1)
				client.CreateImposter(imposter2)

				_, err = client.DeleteAllImposters()
			})
		})

		It("should delete all the installed Imposters at Mountebank", func() {
			impostersURI := MountebankURI + "/imposters"
			resp, _, _ := gorequest.New().Get(impostersURI).End()

			Expect(resp.StatusCode).To(Equal(http.StatusOK))
		})

		It("should not return an error", func() {
			Expect(err).To(BeNil())
		})
	})

})
