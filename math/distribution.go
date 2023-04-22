package math

func Distribution(
	toDivide float64,
	steps int,
) []float64 {
	var result []float64
	seriesSum := float64(SeriesSum(steps))

	for i := 1; i <= steps; i++ {
		a := toDivide * float64(i) / seriesSum
		result = append(result, a)
	}

	return result
}
