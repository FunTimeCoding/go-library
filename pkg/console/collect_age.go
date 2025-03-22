package console

import (
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/time"
)

func CollectAge(v []face.AgeColorable) []float64 {
	var result []float64

	for _, a := range v {
		result = append(result, time.HoursToDecades(a.Age().Hours()))
	}

	return result
}
