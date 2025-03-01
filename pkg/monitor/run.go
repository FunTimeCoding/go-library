package monitor

import (
	"github.com/funtimecoding/go-library/pkg/monitor/report"
	"github.com/funtimecoding/go-library/pkg/notation"
	"github.com/funtimecoding/go-library/pkg/system/run"
)

func Run(name string) *report.Report {
	r := run.New()
	r.Panic = false
	r.Start(name, "--notation")
	result := &report.Report{}
	notation.DecodeStrict(r.OutputString, &result, false)

	return result
}
