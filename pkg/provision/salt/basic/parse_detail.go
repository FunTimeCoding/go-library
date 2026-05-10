package basic

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/web/detail_error"
	"strings"
)

func parseDetail(
	body []byte,
	status string,
) error {
	s := string(body)
	start := strings.Index(s, "<p>")
	end := strings.Index(s, "</p>")

	if start >= 0 && end > start {
		return detail_error.New(s[start+3:end], status)
	}

	return fmt.Errorf("%s", status)
}
