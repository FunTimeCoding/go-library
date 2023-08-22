package math

func Weight(
	a float64,
	b float64,
	aWeight float64,
	bWeight float64,
) float64 {
	sum := aWeight + bWeight

	return ((a * aWeight) + (b * bWeight)) / sum
}
