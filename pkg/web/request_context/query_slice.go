package request_context

import (
	"github.com/funtimecoding/go-library/pkg/strings/split"
)

func (c *Context) QuerySlice(k string) []string {
	return split.Comma(c.Query(k))
}
