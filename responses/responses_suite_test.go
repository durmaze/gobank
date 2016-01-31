package responses_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestResponses(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Responses Builder Test Suite")
}
