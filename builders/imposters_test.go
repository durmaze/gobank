package builders_test

import (
	"encoding/json"

	. "github.com/durmaze/gobank/builders"

	"sync"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Imposter Builder Tests", func() {

	Describe("When building an imposter with protocol, port and a single stub", func() {
		var (
			actualImposterAsMap map[string]interface{}

			expectedProtocol = "http"
			expectedPort     = 4546
			expectedStub     = Stub().Build()

			once sync.Once
		)

		BeforeEach(func() {
			once.Do(func() {
				actualImposter := NewImposterBuilder().Protocol(expectedProtocol).Port(expectedPort).Stubs(expectedStub).Build()

				jsonBytes, _ := json.Marshal(actualImposter)
				actualImposterAsMap = map[string]interface{}{}
				json.Unmarshal(jsonBytes, &actualImposterAsMap)
			})
		})

		It("should create the Imposter with the correct Protocol", func() {
			Expect(actualImposterAsMap).To(HaveKeyWithValue("protocol", expectedProtocol))
		})

		It("should create the Imposter with the correct Port", func() {
			// golang converts numbers to float64 when unmarshalling json to map[string]interface{}
			Expect(actualImposterAsMap).To(HaveKeyWithValue("port", float64(expectedPort)))
		})

		It("should create the Imposter with Stubs", func() {
			Expect(actualImposterAsMap).To(HaveKey("stubs"))
		})

		It("should create the Imposter with one Stub", func() {
			stubs := actualImposterAsMap["stubs"]

			Expect(stubs).To(HaveLen(1))
		})

	})

})
