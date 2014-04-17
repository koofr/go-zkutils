package zkutils_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestZkutils(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Zkutils Suite")
}
