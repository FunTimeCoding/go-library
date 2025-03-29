package workflow

import (
	"github.com/funtimecoding/go-library/pkg/console/status"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
	"github.com/funtimecoding/go-library/pkg/time"
)

func (w *Workflow) Format(f *option.Format) string {
	return status.New(f).String(
		w.Name,
		w.CreatedAt.Format(time.DateMinute),
	).Raw(w).Format()
}
