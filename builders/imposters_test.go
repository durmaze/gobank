package builders_test

import (
	. "github.com/durmaze/gobank/builders"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"sync"
)

var _ = Describe("Imposter Builder Tests", func() {

	Describe("When building an imposter with protocol and port", func() {
		var (
			imposter Imposter
			once sync.Once
		)

		BeforeEach(func() {
			once.Do(func(){
				imposter = NewImposterBuilder().Protocol("http").Port(4546).Build()
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
