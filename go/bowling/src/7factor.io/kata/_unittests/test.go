package _unittests

import (
	. "7factor.io/kata"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func testPlayGame(rolls int, pins int) (int) {
	g := Game{}

	for i := 0; i < rolls; i++ {
		g.Roll(pins)
	}

	return g.CalcScore()
}

func testRollSpare() (int) {
	g := Game{}

	g.Roll(5)
	g.Roll(5)
	g.Roll(3)

	return g.CalcScore()
}

func testRollStrike() (int) {
	g := Game{}

	g.Roll(10)
	g.Roll(3)
	g.Roll(4)

	return g.CalcScore()
}

var _ = Describe("The Bowling Game", func() {
	Context("When running the bowling game", func() {
		It("returns 0 for gutter game", func() {
			score := testPlayGame(20, 0)
			Expect(score).To(Equal(0))
		})

		It("returns 20 for rolling 1 every time", func() {
			score := testPlayGame(20, 1)
			Expect(score).To(Equal(20))
		})

		It("tests one Spare", func() {
			score := testRollSpare()
			Expect(score).To(Equal(16))
		})

		It("tests one Strike", func() {
			score := testRollStrike()
			Expect(score).To(Equal(24))
		})

		It("test perfect game", func() {
			score := testPlayGame(12, 10)
			Expect(score).To(Equal(300))
		})
	})
})
