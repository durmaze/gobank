package imposters_test

import (
	. "github.com/durmaze/gobank/imposters"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"sync"
)

var _ = Describe("When building an imposter with protocol and port", func() {
	var (
		imposter Imposter
		once sync.Once
	)

	BeforeEach(func() {
		once.Do(func(){
			imposter = New().Protocol("http").Port(4546).Build()
		})
  })

  It("should create the Imposter with the specified protocol", func() {
    Expect(imposter.Protocol).To(Equal("http"))
  })

  It("should create the Imposter on the specified port", func() {
    Expect(imposter.Port).To(Equal(4546))
  })

})
