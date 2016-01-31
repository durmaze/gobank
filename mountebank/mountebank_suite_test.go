package mountebank_test

import (
	"net/http"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/parnurzeal/gorequest"
)

var MountebankUri string

func TestMountebank(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Mountebank Integration Test Suite")
}

var _ = BeforeSuite(func() {
	MountebankUri = "http://localhost:2525"

	Expect(isMountebankRunning(MountebankUri)).To(BeTrue(), "Mountebank is not running")

	truncateMountebank(MountebankUri)
})

var _ = AfterSuite(func() {
})

func isMountebankRunning(mountebankBaseUri string) bool {
	resp, _, _ := gorequest.New().Get(mountebankBaseUri).End()

	return resp.StatusCode == http.StatusOK
}

func truncateMountebank(mountebankBaseUri string) {
	impostersUri := mountebankBaseUri + "/imposters"
	gorequest.New().Delete(impostersUri).End()
}