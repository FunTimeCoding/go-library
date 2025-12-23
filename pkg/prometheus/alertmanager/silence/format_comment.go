package silence

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
	"strings"
)

func (s *Silence) formatComment(f *option.Format) string {
	if strings.TrimSpace(strings.ToLower(s.Comment)) == "ack" {
		return ""
	}

	if strings.HasPrefix(
		s.Comment,
		"ACK! This alert was acknowledged using karma on ",
	) {
		return ""
	}

	if f.UseColor {
		return fmt.Sprintf(
			"Comment: %s",
			console.Magenta("%s", s.Comment),
		)
	}

	return s.Comment
}
