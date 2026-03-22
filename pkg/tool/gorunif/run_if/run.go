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
	base := viper.GetString(constant.Base)
	head := viper.GetString(constant.Head)
	path := argument.RequiredPositional(0, "PATH")
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

		if strings.HasPrefix(p, path) {
			changed = true

			if o.Verbose {
				fmt.Printf("Match: %s\n", path)
			}

			break
		}
	}

	if changed {
		run.New().Execute("sh", "-c", execute)
	}
}
