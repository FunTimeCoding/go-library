package face

type ScoreColorable interface {
	Score() float64
	ScoreColorFunction() SprintFunction
	SetScoreColorFunction(SprintFunction)
}
