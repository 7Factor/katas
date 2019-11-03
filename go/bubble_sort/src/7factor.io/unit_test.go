package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"sort"
	"testing"
)

func TestKata(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Golang Kata Unit Test Suite")
}

var _ = Describe("Given an array of ints", func() {
	Context("When passed a nil input", func() {
		It("Should return an empty arr and no errors", func() {
			emptyArr := make([]int, 0)
			actual, err := bubbleSort(nil)
			Expect(actual).To(Equal(emptyArr))
			Expect(err).To(BeNil())
		})
	})

	Context("When passed an empty arr", func() {
		It("Should return an empty arr and no errors", func() {
			emptyArr := make([]int, 0)
			actual, err := bubbleSort(emptyArr)
			Expect(actual).To(Equal(emptyArr))
			Expect(err).To(BeNil())
		})
	})

	Context("When passed an arr with a single element", func() {
		It("Should return the arr and no errors", func() {
			singleElement := []int{1}
			actual, err := bubbleSort(singleElement)
			Expect(actual).To(Equal(singleElement))
			Expect(err).To(BeNil())
		})
	})

	Context("When passed a sorted arr with two elements", func() {
		It("Should return the arr and no errors", func() {
			twoElements := []int{1,2}
			actual, err := bubbleSort(twoElements)
			Expect(actual).To(Equal(twoElements))
			Expect(err).To(BeNil())
		})
	})

	Context("When passed an unsorted arr with two elements", func() {
		It("Should return a sorted arr and no errors", func() {
			unsorted := []int{2,1}
			actual, err := bubbleSort(unsorted)
			Expect(actual).To(Equal([]int{1,2}))
			Expect(err).To(BeNil())
		})
	})

	Context("When passed an unsorted arr with three elements", func() {
		It("Should return the arr and no errors", func() {
			threeElements := []int{1,2,3}
			actual, err := bubbleSort(threeElements)
			Expect(actual).To(Equal(threeElements))
			Expect(err).To(BeNil())
		})
	})

	Context("When passed an unsorted arr with three elements where only the first element is out of order", func() {
		It("Should return a sorted arr and no errors", func() {
			firstUnsorted := []int{2,1,3}
			actual, err := bubbleSort(firstUnsorted)
			Expect(actual).To(Equal([]int{1,2,3}))
			Expect(err).To(BeNil())
		})
	})

	Context("When passed an unsorted arr with three elements where only the second element is out of order", func() {
		It("Should return a sorted arr and no errors", func() {
			secondUnsorted := []int{1,3,2}
			actual, err := bubbleSort(secondUnsorted)
			Expect(actual).To(Equal([]int{1,2,3}))
			Expect(err).To(BeNil())
		})
	})

	Context("When passed an unsorted arr with three elements where the first and second elements are out of order", func() {
		It("Should return a sorted arr and no errors", func() {
			fullyUnsorted := []int{3,2,1}
			actual, err := bubbleSort(fullyUnsorted)
			Expect(actual).To(Equal([]int{1,2,3}))
			Expect(err).To(BeNil())
		})
	})

	Context("When passed a giant input", func() {
		It("Should return a sorted array and no errors", func() {
			bigUnsorted := makeRandomArray(Range{0, 100000})
			actual, err := bubbleSort(bigUnsorted)
			sort.Ints(bigUnsorted)
			Expect(actual).To(Equal(bigUnsorted))
			Expect(err).To(BeNil())
		})
	})
})

type Range struct {
	min int
	max int
}

func makeRandomArray(r Range) []int {
	a := make([]int, r.max-r.min+1)
	for i := range a {
		a[i] = r.min + i
	}
	return a
}
