package channel

import (
	"github.com/funtimecoding/go-library/pkg/console/status"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
)

func (c *Channel) Format(f *option.Format) string {
	return status.New(f).Integer64(c.Identifier).String(c.Name).Format()
}
