package query

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/prometheus/query/filter"
)

func Rate(
	name string,
	f *filter.Filter,
	minutes int,
) string {
	return fmt.Sprintf("rate(%s[%dm])", Filter(name, f), minutes)
}
