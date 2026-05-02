package gofix

func runImportAliasFix(
	patterns []string,
	diff bool,
	r *results,
) {
	runImportAliasFixWithDirectory(patterns, "", diff, r)
}
