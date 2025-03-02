package workflow

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/time"
)

func (w *Workflow) Format() string {
	return fmt.Sprintf(
		"%s | %s",
		w.Name,
		w.CreatedAt.Format(time.DateMinute),
	)
}
