package alert

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/strings/join"
)

func (a *Alert) formatRemainingLabels() string {
	if len(a.RemainingLabels) == 0 {
		return ""
	}

	var elements []string

	for k, v := range a.RemainingLabels {
		elements = append(elements, fmt.Sprintf("%s=%s", k, v))
	}

	return join.Comma(elements)
}
