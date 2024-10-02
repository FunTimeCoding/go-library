package commit

import (
	"github.com/funtimecoding/go-library/pkg/console/format"
	"github.com/funtimecoding/go-library/pkg/console/status"
	"github.com/funtimecoding/go-library/pkg/time"
)

func (c *Commit) Format(s *format.Settings) string {
	return status.New(s).String(
		c.Date.Format(time.DateMinute),
		c.Author,
		c.formatTitle(s),
	).Raw(c.Raw).Format()
}
