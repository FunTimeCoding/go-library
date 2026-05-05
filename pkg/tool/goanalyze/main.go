package goanalyze

import (
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
	"github.com/funtimecoding/go-library/pkg/tool/goanalyze/constant"
	"golang.org/x/tools/go/analysis/multichecker"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	r := reporter.New(constant.Name, version).Start()
	defer func() { r.RecoverFlush(recover()) }()
	multichecker.Main(
		naming.Analyzer,
		forbidden_call.Analyzer,
		forbidden_import.Analyzer,
		string_concatenation.Analyzer,
		string_constant.Analyzer,
		struct_literal.Analyzer,
		call_format.Analyzer,
		defer_close.Analyzer,
		file_identity.Analyzer,
		type_receiver.Analyzer,
		unchecked_print_write.Analyzer,
		anonymous_struct.Analyzer,
	)
}
