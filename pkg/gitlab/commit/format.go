package commit

import (
	"github.com/funtimecoding/go-library/pkg/console/format"
	"github.com/funtimecoding/go-library/pkg/console/status"
	"github.com/funtimecoding/go-library/pkg/time"
)

func (c *Commit) Format(s *format.Settings) string {
	return status.New(s).Integer(c.Project).String(
		c.Identifier,
		c.Date.Format(time.DateMinute),
		c.Author,
		c.Title,
		c.Message,
	).Raw(c.Raw).Format()
}
