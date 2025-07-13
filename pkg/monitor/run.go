package monitor

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/monitor/constant"
	"github.com/funtimecoding/go-library/pkg/monitor/report"
	"github.com/funtimecoding/go-library/pkg/notation"
	"github.com/funtimecoding/go-library/pkg/system/run"
	"log"
	"time"
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
	result := &report.Report{}

	if r.Error != nil {
		s := fmt.Sprintf("run fail: %s %s", name, r.Error)
		log.Print(s)
		result.AddItem(
			fmt.Sprintf("%s-%s", constant.MonitorPrefix, name),
			constant.ErrorLevel,
			s,
			"",
			&time.Time{},
		)

		return result
	}

	if e := notation.Decode(r.OutputString, &result); e != nil {
		log.Printf("parse fail: %s %s", name, e)
		result.AddItem(
			fmt.Sprintf("%s-%s", constant.MonitorPrefix, name),
			constant.ErrorLevel,
			fmt.Sprintf("parse fail: %s %s", name, r.Error),
			"",
			&time.Time{},
		)

		return result
	}

	return result
}
