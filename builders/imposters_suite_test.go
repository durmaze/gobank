package builders_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestImposters(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Mountebank Imposter Builders Test Suite")
}
