package file_report

import (
	"bufio"
	"io"
	"strings"
)

func New(
	path string,
	r io.Reader,
) *Report {
	return &Report{
		FilePath:        path,
		originalBuilder: &strings.Builder{},
		fixedBuilder:    &strings.Builder{},
		scanner:         bufio.NewScanner(r),
	}
}
