package _unittests

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestAPI(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Unit Tests Suite")
}
