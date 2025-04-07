package filter

import "github.com/funtimecoding/go-library/pkg/prometheus/constant"

func (f *Filter) Namespace(s string) *Filter {
	return f.Equal(constant.NamespaceLabel, s)
}
