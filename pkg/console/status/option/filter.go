package option

import "github.com/funtimecoding/go-library/pkg/console/status/option/pair"

func (f *Format) Filter(
	k string,
	v string,
) *Format {
	f.Filters = append(f.Filters, pair.New(k, v))

	return f
}
