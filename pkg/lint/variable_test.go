package lint

import (
	"github.com/funtimecoding/go-library/pkg/lint/concern"
	library "github.com/funtimecoding/go-library/pkg/strings"
	"strings"
	"testing"
)

func TestVariableErrAssignment(t *testing.T) {
	l := Variable(
		library.Alfa,
		strings.NewReader(
			"package example\n\nfunc Example() {\n\terr := foo()\n}\n",
		),
	)
	assertReport(
		t,
		"Alfa",
		true,
		[]*concern.Concern{
			{
				Key:      "err_variable",
				Text:     "Use e instead of err for error variable",
				Path:     "Alfa",
				Line:     4,
				LineText: "\terr := foo()",
			},
		},
		"",
		l,
	)
}

func TestVariableErrMultiReturn(t *testing.T) {
	l := Variable(
		library.Bravo,
		strings.NewReader(
			"package example\n\nfunc Example() {\n\tx, err := foo()\n}\n",
		),
	)
	assertReport(
		t,
		"Bravo",
		true,
		[]*concern.Concern{
			{
				Key:      "err_variable",
				Text:     "Use e instead of err for error variable",
				Path:     "Bravo",
				Line:     4,
				LineText: "\tx, err := foo()",
			},
		},
		"",
		l,
	)
}

func TestVariableErrComparison(t *testing.T) {
	l := Variable(
		library.Charlie,
		strings.NewReader(
			"package example\n\nfunc Example() bool {\n\treturn err == nil\n}\n",
		),
	)
	assertReport(t, "Charlie", false, nil, "", l)
}

func TestVariableEOkay(t *testing.T) {
	l := Variable(
		library.Delta,
		strings.NewReader(
			"package example\n\nfunc Example() {\n\te := foo()\n}\n",
		),
	)
	assertReport(t, "Delta", false, nil, "", l)
}

func TestVariableErrComment(t *testing.T) {
	l := Variable(
		library.Echo,
		strings.NewReader(
			"package example\n\n// err := foo()\n",
		),
	)
	assertReport(t, "Echo", false, nil, "", l)
}
