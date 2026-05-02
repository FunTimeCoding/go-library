package gofix

func runImportAliasFix(
	patterns []string,
	diff bool,
) {
	runImportAliasFixWithDirectory(patterns, "", diff)
}
