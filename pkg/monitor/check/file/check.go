package file

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/monitor/check/file/option"
	"github.com/funtimecoding/go-library/pkg/monitor/constant"
	"github.com/funtimecoding/go-library/pkg/monitor/report"
	"github.com/funtimecoding/go-library/pkg/system"
	"time"
)

func Check(o *option.File) {
	r := report.New()

	if o.Verbose {
		r.AddNote("verbose output")
	}

	for i, p := range collect() {
		t := system.Stat(p).ModTime()

		if time.Since(t) > 5*time.Minute {
			r.AddItem(
				fmt.Sprintf("%s-%d", constant.FilePrefix, i+1),
				constant.WarningLevel,
				fmt.Sprintf("File old: %s", p),
				"",
				&t,
			)
		} else if o.Verbose {
			r.AddItem(
				fmt.Sprintf("%s-%d", constant.FilePrefix, i+1),
				constant.InformationLevel,
				fmt.Sprintf("File good: %s", p),
				"",
				&t,
			)
		} else {
			r.AddNote("good: %s", p)
		}
	}

	r.Print()
}
