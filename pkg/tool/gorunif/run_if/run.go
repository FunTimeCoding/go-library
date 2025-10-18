package run_if

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/strings/split"
	"github.com/funtimecoding/go-library/pkg/system/run"
	"github.com/funtimecoding/go-library/pkg/tool/gorunif/run_if/constant"
	"github.com/funtimecoding/go-library/pkg/tool/gorunif/run_if/option"
	"github.com/spf13/viper"
	"os"
	"os/exec"
	"strings"
)

func Run(o *option.RunIf) {
	// TODO: Build tag format in commit message
	//  Example case: [image:backstage]
	//   If set, build the image
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
		c := exec.Command("sh", "-c", execute)
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
		c.Stdin = os.Stdin
		errors.PanicOnError(c.Run())
	}
}
