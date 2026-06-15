package store

import "fmt"

func orPairs(count int) string {
	var result string

	for i := 0; i < count; i++ {
		if i > 0 {
			result = fmt.Sprintf("%s OR ", result)
		}

		result = fmt.Sprintf("%s(collection = ? AND path = ?)", result)
	}

	return result
}
