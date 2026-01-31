package brew

import (
	"github.com/funtimecoding/go-library/pkg/monitor/report"
	"github.com/funtimecoding/go-library/pkg/system/macos/check/brew/option"
)

func printNotation(o *option.Brew) {
	r := report.New()
	r.AddNote("Option: %+v", o)
	r.Print()
}
