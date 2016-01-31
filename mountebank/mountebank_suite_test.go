package mountebank_test

import (
	"net/http"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/parnurzeal/gorequest"
)

var MountebankBaseUri string

func TestMountebank(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Mountebank Integration Test Suite")
}

var _ = BeforeSuite(func() {
	MountebankBaseUri = "http://localhost:2525"

	Expect(isMountebankRunning(MountebankBaseUri)).To(BeTrue(), "Mountebank is not running")

	truncateMountebank(MountebankBaseUri)
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