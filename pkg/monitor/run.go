package monitor

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/monitor/constant"
	"github.com/funtimecoding/go-library/pkg/monitor/report"
	"github.com/funtimecoding/go-library/pkg/notation"
	"github.com/funtimecoding/go-library/pkg/system/run"
	"log"
)

func Run(name string) *report.Report {
	r := run.New()
	r.Panic = false
	arguments := []string{fmt.Sprintf("--%s", argument.Notation)}

	if false {
		if name == constant.GoFile {
			arguments = append(
				arguments,
				fmt.Sprintf("--%s", argument.Verbose),
			)
		}
	}

	r.Start(append([]string{name}, arguments...)...)

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
