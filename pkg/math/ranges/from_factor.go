package ranges

func FromFactor(
	base float64,
	factor float64,
) Range {
	left := base - (base * factor)
	right := base + (base * factor)

	return Range{L: left, R: right}
}
