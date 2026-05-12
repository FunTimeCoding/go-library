package run

import "github.com/funtimecoding/go-library/pkg/face"

func (r *Run) SetReporter(
	reporter face.Reporter,
	label string,
) {
	r.reporter = reporter
	r.reporterLabel = label
}
