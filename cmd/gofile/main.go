package main

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/monitor"
	"github.com/funtimecoding/go-library/pkg/monitor/constant"
	"github.com/funtimecoding/go-library/pkg/monitor/report"
	"github.com/funtimecoding/go-library/pkg/strings/split"
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/funtimecoding/go-library/pkg/system/environment"
	"github.com/spf13/viper"
	"time"
)

func main() {
	monitor.NotationArgument()
	monitor.VerboseArgument()
	argument.ParseBind()
	r := report.New()
	verbose := viper.GetBool(argument.Verbose)

	if verbose {
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
		} else if verbose {
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

func collect() []string {
	result := argument.Positionals()

	if s := environment.GetDefault(
		"MONITOR_FILE",
		"",
	); s != "" {
		result = append(result, split.Comma(s)...)
	}

	return result
}
