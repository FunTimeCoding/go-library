package generic

import (
	"github.com/funtimecoding/go-library/pkg/prometheus/constant"
	"slices"
)

func sortLabels(s []string) []string {
	var result []string

	if slices.Contains(s, constant.Namespace) {
		result = append(result, constant.Namespace)
	}

	if slices.Contains(s, constant.Pod) {
		result = append(result, constant.Pod)
	}

	if slices.Contains(s, constant.Container) {
		result = append(result, constant.Container)
	}

	for _, e := range s {
		if e == constant.Namespace ||
			e == constant.Pod ||
			e == constant.Container {
			continue
		}

		result = append(result, e)
	}

	return result
}
