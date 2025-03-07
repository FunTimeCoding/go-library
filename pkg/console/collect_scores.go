package console

import "github.com/funtimecoding/go-library/pkg/face"

func CollectScores(v []face.ScoreColorable) []float64 {
	var result []float64

	for _, element := range v {
		if element.Score() != 0 {
			result = append(result, element.Score())
		}
	}

	return result
}
