package monitor

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/monitor/report"
	"github.com/funtimecoding/go-library/pkg/notation"
	"github.com/funtimecoding/go-library/pkg/system/run"
)

func Run(name string) *report.Report {
	r := run.New()
	r.Panic = false
	r.Start(name, fmt.Sprintf("--%s", argument.Notation))
	result := &report.Report{}
	notation.DecodeStrict(r.OutputString, &result, false)

	return result
}
