package strings

func FirstHalf(arguments []string) []string {
	length := len(arguments)

	if length == 0 {
		return []string{}
	}

	count := length / 2

	return arguments[0:count]
}
