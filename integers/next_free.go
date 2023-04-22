package integers

func NextFree(numbers []int) int {
	result := 0

	for {
		if Contains(result, numbers) {
			result++
		} else {
			break
		}
	}

	return result
}
