package gobank_test

import (
	"net/http"
	"os"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/parnurzeal/gorequest"
)

var MountebankURI string

func TestMountebank(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Mountebank Integration Test Suite")
}

var _ = BeforeSuite(func() {
	MountebankURI = os.Getenv("MOUNTEBANK_URI")
	if len(MountebankURI) == 0 {
		MountebankURI = "http://localhost:2525"
	}

	Expect(isMountebankRunning(MountebankURI)).To(BeTrue(), "Mountebank is not running")

	truncateMountebank(MountebankURI)
})

var _ = AfterSuite(func() {
})

func isMountebankRunning(mountebankBaseURI string) bool {
	resp, _, errs := gorequest.New().Get(mountebankBaseURI).End()

	Expect(errs).To(HaveLen(0))

	return resp.StatusCode == http.StatusOK
}

func truncateMountebank(mountebankBaseURI string) {
	impostersURI := mountebankBaseURI + "/imposters"
	gorequest.New().Delete(impostersURI).End()
}
