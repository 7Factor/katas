package kata

type Game struct {
	rolls [21]int
	currentRoll int
}

type Play interface {
	Roll(int)
	CalcScore() int
	sumOfBallsInFrame() int
	spareBonus() int
	strikeBonus() int
	isSpare() bool
}

func (g *Game) Roll(pins int) {
	g.rolls[g.currentRoll] = pins
	g.currentRoll ++
}

func (g *Game) CalcScore() (int) {
	var score int
	var frameIndex int
	for frame := 0; frame < 10; frame++ {
		if g.rolls[frameIndex] == 10 { // strike
			score += 10 + g.strikeBonus(frameIndex)
			frameIndex ++
		} else if g.isSpare(frameIndex) { //spare
			score += 10 + g.spareBonus(frameIndex)
			frameIndex += 2
		} else {
			score += g.sumOfBallsInFrame(frameIndex)
			frameIndex += 2
		}
	}
	return score
}

func (g *Game) sumOfBallsInFrame(frameIndex int) (int) {
	return g.rolls[frameIndex] + g.rolls[frameIndex+1]
}

func (g *Game) spareBonus(frameIndex int) (int) {
	return g.rolls[frameIndex+ 2]
}

func(g *Game) strikeBonus(frameIndex int) (int) {
	return g.rolls[frameIndex+1] + g.rolls[frameIndex+2]
}

func (g *Game) isSpare(frameIndex int) (bool) {
	return g.rolls[frameIndex] + g.rolls[frameIndex+ 1] == 10
}
