package commit

import (
	"github.com/funtimecoding/go-library/pkg/console/status"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
	"github.com/funtimecoding/go-library/pkg/time"
)

func (c *Commit) Format(f *option.Format) string {
	return status.New(f).String(
		c.Date.Format(time.DateMinute),
		c.Author,
		c.formatTitle(f),
	).RawList(c.Raw).Format()
}
