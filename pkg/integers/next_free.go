package integers

func NextFree(
	start int,
	numbers []int,
) int {
	for {
		if Contains(numbers, start) {
			start++
		} else {
			break
		}
	}

	return start
}
