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
	if !all && len(v) > constant.NotationReport {
		v = v[0:constant.NotationReport]
		r.AddItem(
			prefix+"-0",
			constant.WarningLevel,
			fmt.Sprintf(
				"Too many %s (%d), showing only the newest %d",
				name,
				len(v),
				constant.NotationReport,
			),
			"",
			&time.Time{},
		)
	}

	return v
}
