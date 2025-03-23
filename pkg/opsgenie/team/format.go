package team

import (
	"github.com/funtimecoding/go-library/pkg/console/status"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
)

func (t *Team) Format(f *option.Format) string {
	result := status.New(f).String(
		t.Name,
		t.Identifier,
	).Raw(t.Raw)

	if t.Description != "" {
		result.String(t.Description)
	}

	return result.Format()
}
