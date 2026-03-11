package lint

import (
	library "github.com/funtimecoding/go-library/pkg/strings"
	"strings"
	"testing"
)

func TestFunctionClean(t *testing.T) {
	l := Function(
		library.Alfa,
		strings.NewReader(
			"package example\n\nfunc Example() {\n\tfmt.Println(\"hello\")\n}\n",
		),
	)
	assertReport(t, "Alfa", false, nil, "", l)
}
