package constant

const (
	FrontMatterDelimiter = "---"

	FrontMatterDelimiterKey  = "front_matter_delimiter"
	FrontMatterDelimiterText = "No front matter delimiter"

	ImportBlankKey  = "import_blank"
	ImportBlankText = "Import block contains blank line"

	UnsortedImportsKey  = "unsorted_imports"
	UnsortedImportsText = "Imports are not sorted"

	SingleMultiImportKey  = "single_multi_import"
	SingleMultiImportText = "Single multi import"

	EmptyFunctionBodyKey  = "empty-function-body"
	EmptyFunctionBodyText = "Function body with only whitespace"

	ErrVariableKey  = "err_variable"
	ErrVariableText = "Use e instead of err for error variable"

	PackageNameKey  = "package_name"
	PackageNameText = "Blacklisted package name"

	MultipleExportedTypesKey  = "multiple_exported_types"
	MultipleExportedTypesText = "Multiple exported types in one file"

	MissingTestFileKey  = "missing_test_file"
	MissingTestFileText = "No test file in package"
)

var PackageBlocklist = []string{"api"}
