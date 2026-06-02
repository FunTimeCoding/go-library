package convert

import "github.com/funtimecoding/go-library/pkg/prometheus/label_result"

func LabelResult(r *label_result.Result) *SlimLabelResult {
	return &SlimLabelResult{
		Values:   r.Values,
		Warnings: r.Warnings,
	}
}
