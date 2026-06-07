package goaudit

import (
	"github.com/funtimecoding/go-library/pkg/lint/concern"
	"github.com/funtimecoding/go-library/pkg/lint/output"
	"github.com/funtimecoding/go-library/pkg/system/virtual_file_system"
	"github.com/funtimecoding/go-library/pkg/tool/goaudit/scan"
	"os"
)

func runHeadless(
	v *virtual_file_system.System,
	e []*scan.Service,
	identityWarnings []*concern.Concern,
) {
	r := output.NewResults()

	for _, s := range e {
		for _, c := range s.Concerns {
			r.AddConcern(c)
		}
	}

	for _, c := range identityWarnings {
		r.AddConcern(c)
	}

	for _, c := range scan.MissingSentry(v) {
		r.AddConcern(c)
	}

	if output.PrintResults(r.Entries, false) {
		os.Exit(1)
	}
}
