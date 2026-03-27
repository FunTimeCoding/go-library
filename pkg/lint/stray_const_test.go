package lint

import (
	"github.com/funtimecoding/go-library/pkg/lint/concern"
	lintConstant "github.com/funtimecoding/go-library/pkg/lint/constant"
	stringLibrary "github.com/funtimecoding/go-library/pkg/strings"
	"strings"
	"testing"
)

func TestStrayConstantFlagged(t *testing.T) {
	l := StrayConstant(
		stringLibrary.Alfa,
		strings.NewReader(
			"package example\n\nconst Foo = 1\n",
		),
	)
	assertReport(
		t,
		"Alfa",
		true,
		[]*concern.Concern{
			{
				Key:      lintConstant.StrayConstantKey,
				Text:     lintConstant.StrayConstantText,
				Path:     "Alfa",
				Line:     3,
				LineText: "const Foo = 1",
			},
		},
		"",
		l,
	)
}

func TestStrayConstantBlockFlagged(t *testing.T) {
	l := StrayConstant(
		stringLibrary.Bravo,
		strings.NewReader(
			"package example\n\nconst (\n\tFoo = 1\n\tBar = 2\n)\n",
		),
	)
	assertReport(
		t,
		"Bravo",
		true,
		[]*concern.Concern{
			{
				Key:      lintConstant.StrayConstantKey,
				Text:     lintConstant.StrayConstantText,
				Path:     "Bravo",
				Line:     3,
				LineText: "const (",
			},
		},
		"",
		l,
	)
}

func TestStrayConstantExemptByFilename(t *testing.T) {
	l := StrayConstant(
		"constant.go",
		strings.NewReader(
			"package example\n\nconst Foo = 1\n",
		),
	)
	assertReport(t, "constant.go", false, nil, "", l)
}

func TestStrayConstantExemptByFilenameTest(t *testing.T) {
	l := StrayConstant(
		"constant_test.go",
		strings.NewReader(
			"package example\n\nconst Foo = 1\n",
		),
	)
	assertReport(t, "constant_test.go", false, nil, "", l)
}

func TestStrayConstantExemptByPackage(t *testing.T) {
	l := StrayConstant(
		stringLibrary.Charlie,
		strings.NewReader(
			"package constant\n\nconst Foo = 1\n",
		),
	)
	assertReport(t, "Charlie", false, nil, "", l)
}

func TestStrayConstantInsideFunctionNotFlagged(t *testing.T) {
	l := StrayConstant(
		stringLibrary.Delta,
		strings.NewReader(
			"package example\n\nfunc Example() {\n\tconst x = 1\n\t_ = x\n}\n",
		),
	)
	assertReport(t, "Delta", false, nil, "", l)
}
