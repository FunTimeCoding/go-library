package alert

import (
	"fmt"
	"strings"
)

func (a *Alert) formatName() string {
	var result string

	if icons := a.emoji(); len(icons) > 0 {
		result = fmt.Sprintf(
			"%s %s", strings.Join(icons, ""),
			a.Name,
		)
	} else {
		result = a.Name
	}

	return result
}
