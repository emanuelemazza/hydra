package aggregator_test

import (
	. "github.com/innotech/hydra/vendors/github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGinkgoAggregator(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Ginkgo Aggregator Suite")
}
