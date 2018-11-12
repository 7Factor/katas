package _unittests

import (
	"7factor.io/kata"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("The word wrap function", func() {
	Context("When wrapping words", func() {
		It("Should do nothing to an empty string", func() {
			actual := kata.Wrap("", 4)
			Expect(actual).To(Equal(""))
		})

		It("When wrapping words smaller than column length", func() {
			actual := kata.Wrap("word", 5)
			Expect(actual).To(Equal("word"))
		})

		It("When wrapping words greater than column length", func() {
			actual := kata.Wrap("wordword", 4)
			Expect(actual).To(Equal("word\nword"))
		})

		It("When wrapping long columns", func() {
			actual := kata.Wrap("birdbirdbirdistheword", 4)
			Expect(actual).To(Equal("bird\nbird\nbird\nisth\newor\nd"))
		})

		It("When wrapping on a word boundry", func() {
			actual := kata.Wrap("word word", 6)
			Expect(actual).To(Equal("word\nword"))
		})
	})
})