package segment

import (
	"fmt"
	"strings"
)

func FormatMessage(
	applicable []string,
	segment, name string,
) string {
	if len(applicable) == 1 {
		return fmt.Sprintf(
			"use %q instead of %q in %s",
			applicable[0],
			segment,
			name,
		)
	}

	quoted := make([]string, len(applicable))

	for i, a := range applicable {
		quoted[i] = fmt.Sprintf("%q", a)
	}

	return fmt.Sprintf(
		"use %s instead of %q in %s",
		strings.Join(quoted, " or "),
		segment,
		name,
	)
}
