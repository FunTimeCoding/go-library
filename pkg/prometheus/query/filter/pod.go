package filter

import "github.com/funtimecoding/go-library/pkg/prometheus/constant"

func (f *Filter) Pod(s string) *Filter {
	return f.Equal(constant.PodLabel, s)
}
