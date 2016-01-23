package imposters_test

import (
	. "github.com/durmaze/gobank/imposters"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("When building an imposter with protocol and port", func() {
	var imposter Imposter

	BeforeEach(func() {
		imposter = New().Protocol("http").Build()
  })

  It("should create an Imposter with the specified protocol", func() {
    Expect(imposter.Protocol).To(Equal("http"))
  })

})
