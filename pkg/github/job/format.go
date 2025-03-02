package job

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/time"
)

func (j *Job) Format() string {
	return fmt.Sprintf(
		"%s | %s | %s",
		j.Name,
		j.CreatedAt.Format(time.DateMinute),
		j.Raw.GetHeadSHA(),
	)
}
