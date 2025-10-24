package exceptions

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/gw2/constant"
	"github.com/funtimecoding/go-library/pkg/notation"
	"github.com/funtimecoding/go-library/pkg/system"
)

func Parse(path string) []*Exception {
	var result []*Exception
	s := system.ReadFile(path, constant.ExceptionFile)

	if false {
		fmt.Printf("Parsing: %s\n", s)
	}

	notation.DecodeStrict(s, &result, true)

	return result
}
