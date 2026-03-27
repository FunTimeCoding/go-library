package gofix

import (
	"fmt"
	sentry "github.com/funtimecoding/go-library/pkg/errors/sentry/constant"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/go-library/pkg/system/environment"
	"go/token"
	"os"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	if c := environment.Optional(sentry.LocatorEnvironment); c != "" {
		r := reporter.New("gofix", c, "", version)
		r.Start()
		defer func() { r.RecoverFlush(recover()) }()
	}

	diff := false
	var patterns []string

	for _, a := range os.Args[1:] {
		if a == "--diff" {
			diff = true
		} else {
			patterns = append(patterns, a)
		}
	}

	if len(patterns) == 0 {
		fmt.Fprintln(os.Stderr, "usage: gofix [--diff] <packages...>")
		os.Exit(1)
	}

	fileSet := token.NewFileSet()
	all := load(fileSet, patterns)
	violations := findViolations(all)

	if len(violations) == 0 {
		return
	}

	edits := buildAllEdits(all, violations)
	applyEdits(fileSet, edits, diff)
}
