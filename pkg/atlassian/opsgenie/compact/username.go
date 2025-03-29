package compact

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/strings/split/key_value"
)

func Username(s string) string {
	first, last := key_value.Dot(Mail(s))

	return fmt.Sprintf("%c%s", first[0], last)
}
