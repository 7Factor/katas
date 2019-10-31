package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestKata(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Golang Kata Unit Test Suite")
}

var _ = Describe("Given an array of ints", func() {
	Context("When passed an empty array", func() {
		It("Should return an empty array and no errors.", func() {
			var emptyArr []int
			actual, err := BubbleSort(emptyArr)
			Expect(actual).To(Equal(emptyArr))
			Expect(err).To(BeNil())
		})
	})
})
