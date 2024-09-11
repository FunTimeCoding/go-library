package task

import (
	"github.com/funtimecoding/go-library/pkg/console/format"
	"github.com/funtimecoding/go-library/pkg/console/status"
)

func (t *Task) Format(s *format.Settings) string {
	return status.New(s).String(string(t.Type), t.State, t.Body).Format()
}
