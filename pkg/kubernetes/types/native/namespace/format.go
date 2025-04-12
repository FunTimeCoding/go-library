package namespace

import (
	"github.com/funtimecoding/go-library/pkg/console/status"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
)

func (n *Namespace) Format(f *option.Format) string {
	return status.New(f).String(
		n.Name,
		n.Cluster,
	).RawList(n.Raw).Format()
}
