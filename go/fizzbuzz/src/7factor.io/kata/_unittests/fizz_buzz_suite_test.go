package _unittests

import (
	"7factor.io/kata"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("The fizzbuzz function", func() {
	Context("When running FizzBuzz", func() {
		It("Should return a list of the correct size", func() {
			actual := kata.FizzBuzz(1, 100)
			Expect(len(actual)).To(Equal(100))
		})

		It("Should print the correct number for every non-multiple of 3", func() {
			var i int
			actual := kata.FizzBuzz(1, 100)
			Expect(actual[i]).To(Equal(kata.Output(i + 1)))
		})

		It("Should print `Fizz` for every number that is a multiple of 3", func() {
			var i int
			actual := kata.FizzBuzz(1, 100)
			Expect(actual[i + 2]).To(Equal("Fizz"))
		})

		It("Should print `Buzz` for every number that is a multiple of 5", func() {
			var i int
			actual := kata.FizzBuzz(1, 100)
			Expect(actual[i + 4]).To(Equal("Buzz"))
		})

		It("Should print `FizzBuzz` for every number that is a multiple of 3 and 5", func() {
			var i int
			actual := kata.FizzBuzz(1, 100)
			Expect(actual[i + 14]).To(Equal("FizzBuzz"))
		})
	})
})