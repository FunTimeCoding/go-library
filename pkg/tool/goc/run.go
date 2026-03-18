package goc

import (
	"github.com/funtimecoding/go-library/pkg/ceph/goc"
	"github.com/funtimecoding/go-library/pkg/tool/goc/option"
)

func Run(o *option.Ceph) {
	goc.Run(o.Command, false)
}
