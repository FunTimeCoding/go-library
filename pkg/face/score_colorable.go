package face

type ScoreColorable interface {
	Score() float64
	ScoreColor() SprintFunction
	SetScoreColor(SprintFunction)
}
