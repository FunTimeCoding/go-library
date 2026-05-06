package run_if

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/strings/split"
	"github.com/funtimecoding/go-library/pkg/system/run"
	"github.com/funtimecoding/go-library/pkg/tool/gorunif/run_if/constant"
	"github.com/funtimecoding/go-library/pkg/tool/gorunif/run_if/option"
	"github.com/spf13/viper"
	"strings"
)

func Run(o *option.RunIf) {
	base := resolveBase(viper.GetString(constant.Base))
	head := viper.GetString(constant.Head)
	suffix := viper.GetBool(constant.Suffix)
	pattern := argument.RequiredPositional(0, "PATTERN")
	execute := argument.RequiredPositional(1, "EXECUTE")
	r := run.New()
	r.Start(
		"git",
		"diff",
		"--name-only",
		fmt.Sprintf("%s..%s", base, head),
	)
	changed := false

	for _, p := range split.NewLine(strings.TrimSpace(r.OutputString)) {
		if o.Verbose {
			fmt.Printf("Change: %s\n", p)
		}

		var match bool

		if suffix {
			match = strings.HasSuffix(p, pattern)
		} else {
			match = strings.HasPrefix(p, pattern)
		}

		if match {
			changed = true

			if o.Verbose {
				fmt.Printf("Match: %s\n", pattern)
			}

			break
		}
	}

	if changed {
		run.New().Execute("sh", "-c", execute)
	}
}
