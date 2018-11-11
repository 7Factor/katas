package _unittests

import (
	"testing"
	. "github.com/onsi/gomega"
	. "github.com/onsi/ginkgo"
)

func TestAPI(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Unit Tests Suite")
}
