package code

import (
	"github.com/funtimecoding/go-library/pkg/console/status"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
)

func (c *Code) Format(f *option.Format) string {
	return status.New(f).String(
		c.Hash,
		c.Name,
		c.Path,
	).RawList(c.Raw).Format()
}
