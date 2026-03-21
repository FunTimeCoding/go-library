package lint

import (
	"github.com/funtimecoding/go-library/pkg/lint/concern"
	lintConstant "github.com/funtimecoding/go-library/pkg/lint/constant"
	stringLibrary "github.com/funtimecoding/go-library/pkg/strings"
	"strings"
	"testing"
)

func TestFunctionCountMultiple(t *testing.T) {
	l := FunctionCount(
		stringLibrary.Alfa,
		strings.NewReader(
			"package example\n\nfunc First() {}\n\nfunc Second() {}\n",
		),
	)
	assertReport(
		t,
		"Alfa",
		true,
		[]*concern.Concern{
			{
				Key:      lintConstant.MultipleFunctionsKey,
				Text:     lintConstant.MultipleFunctionsText,
				Path:     "Alfa",
				Line:     3,
				LineText: "func First() {}",
			},
		},
		"",
		l,
	)
}

func TestFunctionCountSingle(t *testing.T) {
	l := FunctionCount(
		stringLibrary.Bravo,
		strings.NewReader(
			"package example\n\nfunc Only() {}\n",
		),
	)
	assertReport(t, "Bravo", false, nil, "", l)
}

func TestFunctionCountTestFileExempt(t *testing.T) {
	l := FunctionCount(
		"example_test.go",
		strings.NewReader(
			"package example\n\nfunc TestOne(t *testing.T) {}\n\nfunc TestTwo(t *testing.T) {}\n",
		),
	)
	assertReport(t, "example_test.go", false, nil, "", l)
}

func TestFunctionCountMethod(t *testing.T) {
	l := FunctionCount(
		stringLibrary.Charlie,
		strings.NewReader(
			"package example\n\ntype T struct{}\n\nfunc (t *T) First() {}\n\nfunc (t *T) Second() {}\n",
		),
	)
	assertReport(
		t,
		"Charlie",
		true,
		[]*concern.Concern{
			{
				Key:      lintConstant.MultipleFunctionsKey,
				Text:     lintConstant.MultipleFunctionsText,
				Path:     "Charlie",
				Line:     5,
				LineText: "func (t *T) First() {}",
			},
		},
		"",
		l,
	)
}

func TestFunctionCountFuncLiteralNotCounted(t *testing.T) {
	l := FunctionCount(
		stringLibrary.Delta,
		strings.NewReader(
			"package example\n\nfunc Only() {\n\tf := func() {}\n\tf()\n}\n",
		),
	)
	assertReport(t, "Delta", false, nil, "", l)
}
