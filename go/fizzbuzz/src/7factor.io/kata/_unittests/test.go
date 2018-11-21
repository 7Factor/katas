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

		It("Should print the correct number for every non-multiple of 3 and 5", func() {
			var i int
			actual := kata.FizzBuzz(1, 100)
			Expect(actual[i]).To(Equal(kata.Output(i + 1)))
			Expect(actual[i + 99]).To(Equal(kata.Output(i + 100)))
		})

		It("Should print `Fizz` for every number that is a multiple of 3 but not 5", func() {
			var i int
			actual := kata.FizzBuzz(1, 100)
			Expect(actual[i + 2]).To(Equal("Fizz"))
			Expect(actual[i + 98]).To(Equal("Fizz"))
		})

		It("Should print `Buzz` for every number that is a multiple of 5 but not 3", func() {
			var i int
			actual := kata.FizzBuzz(1, 100)
			Expect(actual[i + 4]).To(Equal("Buzz"))
			Expect(actual[i + 94]).To(Equal("Buzz"))
		})

		It("Should print `FizzBuzz` for every number that is a multiple of 3 and 5", func() {
			var i int
			actual := kata.FizzBuzz(1, 100)
			Expect(actual[i + 14]).To(Equal("FizzBuzz"))
			Expect(actual[i + 59]).To(Equal("FizzBuzz"))
		})
	})
})