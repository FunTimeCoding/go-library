package strings

import "github.com/funtimecoding/go-library/pkg/math/normalize"

func Intensity(
	elements []string,
	intensity float64,
) string {
	normalize.Float(&intensity, 0, 1)
	length := len(elements)
	index := int(float64(length) * intensity)

	if index == length {
		index--
	}

	return elements[index]
}
