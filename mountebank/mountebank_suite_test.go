package mountebank_test

import (
	"net/http"
	"os"
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
	MountebankUri = os.Getenv("MOUNTEBANK_URI")
	if len(MountebankUri) == 0 {
		MountebankUri = "http://localhost:2525"
	}

	Expect(isMountebankRunning(MountebankUri)).To(BeTrue(), "Mountebank is not running")

	truncateMountebank(MountebankUri)
})

var _ = AfterSuite(func() {
})

func isMountebankRunning(mountebankBaseUri string) bool {
	resp, _, errs := gorequest.New().Get(mountebankBaseUri).End()

	Expect(errs).To(HaveLen(0))

	return resp.StatusCode == http.StatusOK
}

func truncateMountebank(mountebankBaseUri string) {
	impostersUri := mountebankBaseUri + "/imposters"
	gorequest.New().Delete(impostersUri).End()
}
