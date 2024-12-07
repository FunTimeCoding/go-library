package meta

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/notation"
	"github.com/funtimecoding/go-library/pkg/system"
)

func Parse(path string) *Meta {
	var result *Meta
	s := system.ReadFile(path)

	if false {
		fmt.Printf("Parsing: %s\n", s)
	}

	notation.DecodeStrict(s, &result, true)

	return result
}
