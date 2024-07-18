package status

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/strings/join"
	"strings"
)

func (s *Status) Format() string {
	result := strings.Join(s.bubbles, " | ")

	if s.extended && len(s.lines) > 0 {
		result = fmt.Sprintf("%s\n%s", result, join.NewLine(s.lines))
	}

	return result
}
