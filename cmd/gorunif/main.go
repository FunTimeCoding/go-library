package main

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/monitor"
	"github.com/funtimecoding/go-library/pkg/strings/split"
	"github.com/funtimecoding/go-library/pkg/system/run"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"os"
	"os/exec"
	"strings"
)

const (
	Base = "base"
	Head = "head"
)

func main() {
	pflag.String(Base, "HEAD~1", "Base commit")
	pflag.String(Head, "HEAD", "Head commit")
	monitor.VerboseArgument()
	argument.ParseBind()
	base := viper.GetString(Base)
	head := viper.GetString(Head)
	path := argument.RequiredPositional(0, "PATH")
	execute := argument.RequiredPositional(1, "EXECUTE")
	verbose := viper.GetBool(argument.Verbose)
	r := run.New()
	r.Start(
		"git",
		"diff",
		"--name-only",
		fmt.Sprintf("%s..%s", base, head),
	)
	changed := false

	for _, p := range split.NewLine(strings.TrimSpace(r.OutputString)) {
		if verbose {
			fmt.Printf("Change: %s\n", p)
		}

		if strings.HasPrefix(p, path) {
			changed = true

			if verbose {
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
