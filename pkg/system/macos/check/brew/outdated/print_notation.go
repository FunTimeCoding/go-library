package outdated

import (
	"github.com/funtimecoding/go-library/pkg/monitor/constant"
	item "github.com/funtimecoding/go-library/pkg/monitor/item/constant"
	"github.com/funtimecoding/go-library/pkg/monitor/report"
	"github.com/funtimecoding/go-library/pkg/system/macos/brew/formula"
	"github.com/funtimecoding/go-library/pkg/system/macos/check/brew/outdated/option"
)

func printNotation(
	v []*formula.Formula,
	o *option.Outdated,
) {
	r := report.New()

	for _, e := range report.Trim(
		v,
		r,
		o.All,
		item.GoBrew,
	) {
		r.AddItem(
			item.GoBrew,
			e.MonitorIdentifier,
			constant.Warning,
			e.CurrentVersion,
			e.Link,
			nil,
		)
	}

	r.Print()
}
