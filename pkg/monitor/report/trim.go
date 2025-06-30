package report

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/monitor/constant"
	"time"
)

func Trim[T any](
	v []T,
	r *Report,
	all bool,
	name string,
	prefix string,
) []T {
	count := len(v)

	if !all && count > constant.NotationReport {
		v = v[0:constant.NotationReport]
		r.AddItem(
			prefix+"-0",
			constant.WarningLevel,
			fmt.Sprintf(
				"Too many %s (%d), showing only the newest %d",
				name,
				count,
				constant.NotationReport,
			),
			"",
			&time.Time{},
		)
	}

	return v
}
