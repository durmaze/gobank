package mountebank_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/parnurzeal/gorequest"
	"testing"
	"net/http"
)


func TestMountebank(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Mountebank Test Suite")
}

var _ = BeforeSuite(func() {
  Expect(isMountebankRunning()).To(BeTrue())
})

var _ = AfterSuite(func() {
})

func isMountebankRunning() bool {
	resp, _, _ := gorequest.New().Get("http://localhost:2525").End()

	return resp.StatusCode == http.StatusOK
}