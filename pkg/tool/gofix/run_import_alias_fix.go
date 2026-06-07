package gofix

import "github.com/funtimecoding/go-library/pkg/lint/output"

func runImportAliasFix(
	patterns []string,
	diff bool,
	r *output.Results,
) {
	runImportAliasFixWithDirectory(patterns, "", diff, r)
}
