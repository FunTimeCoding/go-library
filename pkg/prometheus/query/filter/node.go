package filter

import "github.com/funtimecoding/go-library/pkg/prometheus/constant"

func (f *Filter) Node(s string) *Filter {
	return f.Equal(constant.Node, s)
}
