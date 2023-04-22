package normalize

func Integer(
	value *int,
	minimum int,
	maximum int,
) {
	if *value < minimum {
		*value = minimum
	} else if maximum > minimum && *value > maximum {
		*value = maximum
	}
}
