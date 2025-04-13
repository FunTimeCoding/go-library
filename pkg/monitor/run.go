package monitor

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/monitor/report"
	"github.com/funtimecoding/go-library/pkg/notation"
	"github.com/funtimecoding/go-library/pkg/system/run"
	"log"
)

func Run(name string) *report.Report {
	r := run.New()
	r.Panic = false
	r.Start(name, fmt.Sprintf("--%s", argument.Notation))

	if r.Error != nil {
		log.Printf("run fail: %s %s", name, r.Error)

		return nil
	}

	result := &report.Report{}

	if e := notation.Decode(r.OutputString, &result); e != nil {
		log.Printf("parse fail: %s %s", name, e)

		return nil
	}

	return result
}
