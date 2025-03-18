package filter

import "github.com/funtimecoding/go-library/pkg/prometheus/constant"

func (f *Filter) Instance(s string) *Filter {
	return f.Equal(constant.Instance, s)
}
