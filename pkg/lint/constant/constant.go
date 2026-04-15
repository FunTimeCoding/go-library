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

	ErrorVariableKey  = "err_variable"
	ErrorVariableText = "Use e instead of err for error variable"

	PackageNameKey  = "package_name"
	PackageNameText = "Blacklisted package name"

	MultipleExportedTypesKey  = "multiple_exported_types"
	MultipleExportedTypesText = "Multiple exported types in one file"

	MissingTestFileKey  = "missing_test_file"
	MissingTestFileText = "No test file in package"

	MissingBlankBeforeControlKey  = "missing_blank_before_control"
	MissingBlankBeforeControlText = "Missing blank line before control block"

	MissingBlankBeforeExitKey  = "missing_blank_before_exit"
	MissingBlankBeforeExitText = "Missing blank line before return/break/continue"

	ExtraneousBlankLineKey  = "extraneous_blank_line"
	ExtraneousBlankLineText = "Extraneous blank line"

	MissingBlankAfterControlKey  = "missing_blank_after_control"
	MissingBlankAfterControlText = "Missing blank line after control block"

	MissingSentryKey  = "missing_sentry"
	MissingSentryText = "No sentry reporter in program"

	ForbiddenImportKey  = "forbidden_import"
	ForbiddenImportText = `Use "github.com/spf13/pflag" instead of "flag"`

	MultipleFunctionsKey  = "multiple_functions"
	MultipleFunctionsText = "Multiple function definitions in one file"

	StrayConstantKey  = "stray_const"
	StrayConstantText = "Top-level const outside of a constant file or constant package"

	BlankInsideFunctionKey  = "blank_inside_function"
	BlankInsideFunctionText = "Blank line between statements inside function body"

	MissingBlankBeforeDeclarationKey  = "missing_blank_before_declaration"
	MissingBlankBeforeDeclarationText = "Missing blank line before declaration"

	ExtraneousTopLevelBlankKey  = "extraneous_top_level_blank"
	ExtraneousTopLevelBlankText = "Extraneous blank line between top-level declarations"

	MissingBlankBetweenVariableConstantKey  = "missing_blank_between_var_const"
	MissingBlankBetweenVariableConstantText = "Missing blank line between top-level const and var declarations"

	TrackedBinaryKey  = "tracked_binary"
	TrackedBinaryText = "Tracked binary executable"

	DeferCloseKey  = "defer_close"
	DeferCloseText = "Use defer errors.PanicClose instead of defer x.Close()"

	ErrorsImportPath = `"github.com/funtimecoding/go-library/pkg/errors"`
)

var PackageBlocklist = []string{"api"}
