package lint

import (
	"github.com/funtimecoding/go-library/pkg/lint/concern"
	lintConstant "github.com/funtimecoding/go-library/pkg/lint/constant"
	stringLibrary "github.com/funtimecoding/go-library/pkg/strings"
	"strings"
	"testing"
)

func TestStrayConstFlagged(t *testing.T) {
	l := StrayConst(
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
				Key:      lintConstant.StrayConstKey,
				Text:     lintConstant.StrayConstText,
				Path:     "Alfa",
				Line:     3,
				LineText: "const Foo = 1",
			},
		},
		"",
		l,
	)
}

func TestStrayConstBlockFlagged(t *testing.T) {
	l := StrayConst(
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
				Key:      lintConstant.StrayConstKey,
				Text:     lintConstant.StrayConstText,
				Path:     "Bravo",
				Line:     3,
				LineText: "const (",
			},
		},
		"",
		l,
	)
}

func TestStrayConstExemptByFilename(t *testing.T) {
	l := StrayConst(
		"constant.go",
		strings.NewReader(
			"package example\n\nconst Foo = 1\n",
		),
	)
	assertReport(t, "constant.go", false, nil, "", l)
}

func TestStrayConstExemptByFilenameTest(t *testing.T) {
	l := StrayConst(
		"constant_test.go",
		strings.NewReader(
			"package example\n\nconst Foo = 1\n",
		),
	)
	assertReport(t, "constant_test.go", false, nil, "", l)
}

func TestStrayConstExemptByPackage(t *testing.T) {
	l := StrayConst(
		stringLibrary.Charlie,
		strings.NewReader(
			"package constant\n\nconst Foo = 1\n",
		),
	)
	assertReport(t, "Charlie", false, nil, "", l)
}

func TestStrayConstInsideFunctionNotFlagged(t *testing.T) {
	l := StrayConst(
		stringLibrary.Delta,
		strings.NewReader(
			"package example\n\nfunc Example() {\n\tconst x = 1\n\t_ = x\n}\n",
		),
	)
	assertReport(t, "Delta", false, nil, "", l)
}
