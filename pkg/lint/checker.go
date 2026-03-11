package lint

import (
	"github.com/funtimecoding/go-library/pkg/lint/file_report"
	"io"
)

type Checker func(
	string,
	io.Reader,
) *file_report.Report
