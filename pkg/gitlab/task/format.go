package task

import (
	"github.com/funtimecoding/go-library/pkg/console/format"
	"github.com/funtimecoding/go-library/pkg/console/status"
)

func (t *Task) Format(f *format.Settings) string {
	return status.New(f).String(
		string(t.Type),
		t.State,
		t.Body,
	).Raw(t.Raw).Format()
}
