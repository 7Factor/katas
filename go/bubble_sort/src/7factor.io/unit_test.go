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
	Context("When passed a nil input", func() {
		It("Should return an empty array and no errors", func() {
			emptyArr := make([]int, 0)
			actual, err := bubbleSort(nil)
			Expect(actual).To(Equal(emptyArr))
			Expect(err).To(BeNil())
		})
	})

	Context("When passed an empty array", func() {
		It("Should return an empty array and no errors", func() {
			emptyArr := make([]int, 0)
			actual, err := bubbleSort(emptyArr)
			Expect(actual).To(Equal(emptyArr))
			Expect(err).To(BeNil())
		})
	})

	Context("When passed an array with a single element", func() {
		It("Should return the array and no errors", func() {
			singleElement := []int{1}
			actual, err := bubbleSort(singleElement)
			Expect(actual).To(Equal(singleElement))
			Expect(err).To(BeNil())
		})
	})

	Context("When passed a sorted array with two elements", func() {
		It("Should return the array and no errors", func() {
			twoElements := []int{1,2}
			actual, err := bubbleSort(twoElements)
			Expect(actual).To(Equal(twoElements))
			Expect(err).To(BeNil())
		})
	})
})
