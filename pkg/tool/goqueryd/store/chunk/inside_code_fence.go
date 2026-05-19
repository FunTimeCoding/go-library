package chunk

func insideCodeFence(
	position int,
	fences []codeFence,
) bool {
	for _, f := range fences {
		if position > f.start && position < f.end {
			return true
		}
	}

	return false
}
