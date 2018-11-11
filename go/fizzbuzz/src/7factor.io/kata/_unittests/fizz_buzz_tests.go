package _unittests

import (
	"7factor.io/kata"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("The fizzbuzz function", func() {
	Context("When printing numbers", func() {
		It("Should print out the numbers from 1 to 100", func() {
			actual := kata.FizzBuzz(100)
			Expect(actual).To(Equal(100))
		})
	})
})