package query

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/prometheus/query/filter"
	"github.com/funtimecoding/go-library/pkg/strings/join"
)

func SumByRate(
	name string,
	f *filter.Filter,
	minutes int,
	by ...string,
) string {
	return fmt.Sprintf(
		"sum(%s) by (%s)",
		Rate(name, f, minutes),
		join.CommaSpace(by),
	)
}
