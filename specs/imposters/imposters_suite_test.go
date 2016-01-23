package imposters_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestImposters(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Imposters Test Suite")
}
