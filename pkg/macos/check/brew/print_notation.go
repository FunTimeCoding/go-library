package brew

import (
	"github.com/funtimecoding/go-library/pkg/macos/check/brew/option"
	"github.com/funtimecoding/go-library/pkg/monitor/report"
)

func printNotation(o *option.Brew) {
	r := report.New()
	r.AddNote("Option: %+v", o)
	r.Print()
}
