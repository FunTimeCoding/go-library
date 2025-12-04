package pull_request

import (
	"github.com/funtimecoding/go-library/pkg/console/status"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
	"github.com/funtimecoding/go-library/pkg/time"
)

func (r *Request) Format(f *option.Format) string {
	return status.New(f).String(
		r.formatName(f),
		r.Create.Format(time.DateMinute),
	).RawList(r.Raw).Format()
}
