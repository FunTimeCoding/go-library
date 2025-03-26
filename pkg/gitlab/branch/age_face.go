package branch

import "github.com/funtimecoding/go-library/pkg/face"

func AgeFace(v []*Branch) []face.AgeColorable {
	var result []face.AgeColorable

	for _, e := range v {
		result = append(result, e)
	}

	return result
}
