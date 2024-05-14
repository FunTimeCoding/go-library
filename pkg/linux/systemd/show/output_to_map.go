package show

import (
	"github.com/funtimecoding/go-library/pkg/strings/split"
	"github.com/funtimecoding/go-library/pkg/strings/split/key_value"
)

func OutputToMap(s string) map[string]string {
	result := make(map[string]string)

	for _, element := range split.NewLine(s) {
		k, v := key_value.Equals(element)
		result[k] = v
	}

	return result
}
