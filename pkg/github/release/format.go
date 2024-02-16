package release

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/time"
)

func (r *Release) Format() string {
	return fmt.Sprintf(
		"%s | %s",
		r.Name,
		r.CreatedAt.Format(time.DateMinute),
	)
}
