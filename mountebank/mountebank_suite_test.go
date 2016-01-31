package mountebank_test

import (
	"net/http"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/parnurzeal/gorequest"
)

func TestMountebank(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Mountebank Integration Test Suite")
}

var _ = BeforeSuite(func() {
	Expect(isMountebankRunning()).To(BeTrue(), "Mountebank is not running")
})

var _ = AfterSuite(func() {
})

func isMountebankRunning() bool {
	resp, _, _ := gorequest.New().Get("http://localhost:2525").End()

	return resp.StatusCode == http.StatusOK
}
