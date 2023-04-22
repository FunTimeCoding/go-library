package strings

func SecondHalf(arguments []string) []string {
	length := len(arguments)

	if length == 0 {
		return []string{}
	}

	count := length / 2

	return arguments[count:]
}
