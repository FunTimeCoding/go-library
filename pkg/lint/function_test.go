package lint

import (
	"github.com/funtimecoding/go-library/pkg/lint/concern"
	library "github.com/funtimecoding/go-library/pkg/strings"
	"strings"
	"testing"
)

func TestFunctionEmptyBody(t *testing.T) {
	l := Function(
		library.Bravo,
		strings.NewReader(
			"package main\n\nfunc main() {\n}\n",
		),
	)
	assertReport(
		t,
		"Bravo",
		true,
		[]*concern.Concern{
			{
				Key:      "empty-function-body",
				Text:     "Function body with only whitespace",
				Path:     "Bravo",
				Line:     3,
				LineText: "func main() {\n}",
				Fixed:    true,
			},
		},
		"package main\n\nfunc main() {}\n",
		l,
	)
}

func TestFunctionClean(t *testing.T) {
	l := Function(
		library.Alfa,
		strings.NewReader(
			"package example\n\nfunc Example() {\n\tfmt.Println(\"hello\")\n}\n",
		),
	)
	assertReport(t, "Alfa", false, nil, "", l)
}
