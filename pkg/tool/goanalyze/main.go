package goanalyze

import (
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/go-library/pkg/lint/analyzer/anonymous_struct"
	"github.com/funtimecoding/go-library/pkg/lint/analyzer/call_format"
	"github.com/funtimecoding/go-library/pkg/lint/analyzer/defer_close"
	"github.com/funtimecoding/go-library/pkg/lint/analyzer/file_identity"
	"github.com/funtimecoding/go-library/pkg/lint/analyzer/forbidden_call"
	"github.com/funtimecoding/go-library/pkg/lint/analyzer/forbidden_import"
	"github.com/funtimecoding/go-library/pkg/lint/analyzer/naming"
	"github.com/funtimecoding/go-library/pkg/lint/analyzer/string_concatenation"
	"github.com/funtimecoding/go-library/pkg/lint/analyzer/string_constant"
	"github.com/funtimecoding/go-library/pkg/lint/analyzer/struct_literal"
	"github.com/funtimecoding/go-library/pkg/lint/analyzer/type_receiver"
	"github.com/funtimecoding/go-library/pkg/lint/analyzer/unchecked_print_write"
	"github.com/funtimecoding/go-library/pkg/lint/output"
	"github.com/funtimecoding/go-library/pkg/tool/goanalyze/constant"
	"go/token"
	"golang.org/x/tools/go/packages"
	"os"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	r := reporter.New(constant.Identity.Name(), version).Start()
	defer func() { r.RecoverFlush(recover()) }()
	a := argument.NewInstance(constant.Identity)
	a.Boolean("summary", false, "One line per file")
	a.Parse(version, gitHash, buildDate)
	patterns := a.Positionals()

	if len(patterns) == 0 {
		patterns = []string{"./..."}
	}

	fileSet := token.NewFileSet()
	loaded := load(fileSet, patterns)
	results := output.NewResults()

	for _, p := range loaded {
		run(p, &results)
	}

	hasBlocked := output.PrintResults(results.Entries, a.GetBoolean("summary"))

	if hasBlocked {
		os.Exit(1)
	}
}

func run(
	p *packages.Package,
	results *output.Results,
) {
	naming.Check(p, results)
	forbidden_call.Check(p, results)
	forbidden_import.Check(p, results)
	string_concatenation.Check(p, results)
	string_constant.Check(p, results)
	struct_literal.Check(p, results)
	call_format.Check(p, results)
	defer_close.Check(p, results)
	file_identity.Check(p, results)
	type_receiver.Check(p, results)
	unchecked_print_write.Check(p, results)
	anonymous_struct.Check(p, results)
}
