package reporters_test

import (
	. "github.com/innotech/hydra/vendors/github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestReporters(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Reporters Suite")
}
