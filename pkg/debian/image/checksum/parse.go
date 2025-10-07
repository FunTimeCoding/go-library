package checksum

import (
	"github.com/funtimecoding/go-library/pkg/strings/separator"
	"github.com/funtimecoding/go-library/pkg/strings/split"
	"strings"
)

func Parse(input string) map[string]string {
	result := make(map[string]string)

	for _, element := range split.NewLine(input) {
		parts := strings.Split(element, separator.DoubleSpace)
		result[parts[1]] = parts[0]
	}

	return result
}
