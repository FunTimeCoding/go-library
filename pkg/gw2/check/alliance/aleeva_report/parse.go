package aleeva_report

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/notation"
	"github.com/funtimecoding/go-library/pkg/system"
)

func Parse(base string, name string) []*Report {
	var result []*Report
	s := system.ReadFile(base, name)

	if false {
		fmt.Printf("Parsing: %s\n", s)
	}

	notation.DecodeStrict(s, &result, false)

	return result
}
