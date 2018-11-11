package _unittests

import (
	"7factor.io/kata"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"strconv"
)

var _ = Describe("The fizzbuzz function", func() {
	Context("When printing outputs", func() {
		It("Should return a list of the correct size", func() {
			actual := kata.FizzBuzz(1, 100)
			Expect(len(actual)).To(Equal(100))
		})

		// TODO: test all values, not just hard coded ones
		It("Should print the correct number for every non-multiple of 3", func() {
			actual := kata.FizzBuzz(1, 100)
			Expect(actual[0]).To(Equal(strconv.Itoa(1)))
			Expect(actual[1]).To(Equal(strconv.Itoa(2)))
		})

		It("Should print `Fizz` for every number that is a multiple of 3", func() {
			actual := kata.FizzBuzz(1, 100)
			Expect(actual[2]).To(Equal("Fizz"))
		})

		It("Should print `Buzz` for every number that is a multiple of 5", func() {
			actual := kata.FizzBuzz(1, 100)
			Expect(actual[4]).To(Equal("Buzz"))
		})

		It("Should print `FizzBuzz` for every number that is a multiple of 3 and 5", func() {
			actual := kata.FizzBuzz(1, 100)
			Expect(actual[14]).To(Equal("FizzBuzz"))
		})
	})
})