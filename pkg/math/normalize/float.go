package normalize

func Float(
	value *float64,
	minimum float64,
	maximum float64,
) {
	if *value < minimum {
		*value = minimum
	} else if maximum > minimum && *value > maximum {
		*value = maximum
	}
}
