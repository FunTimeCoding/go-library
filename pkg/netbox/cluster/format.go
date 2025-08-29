package cluster

import (
	"github.com/funtimecoding/go-library/pkg/console/status"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
)

func (c *Cluster) Format(f *option.Format) string {
	return status.New(f).String(c.Name).RawList(c.Raw).Format()
}
