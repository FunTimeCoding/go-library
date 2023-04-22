package math

func DecayRate(time float64) float64 {
	return Logarithm(0.5) / -time
}
