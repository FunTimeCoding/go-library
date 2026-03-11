package lint

import (
	"github.com/funtimecoding/go-library/pkg/lint/concern"
	library "github.com/funtimecoding/go-library/pkg/strings"
	"strings"
	"testing"
)

func TestTypeCountMultiple(t *testing.T) {
	l := TypeCount(
		library.Alfa,
		strings.NewReader(
			"package example\n\ntype Foo struct{}\n\ntype Bar struct{}\n",
		),
	)
	assertReport(
		t,
		"Alfa",
		true,
		[]*concern.Concern{
			{
				Key:      "multiple_exported_types",
				Text:     "Multiple exported types in one file",
				Path:     "Alfa",
				Line:     3,
				LineText: "type Foo struct{}",
			},
		},
		"",
		l,
	)
}

func TestTypeCountSingle(t *testing.T) {
	l := TypeCount(
		library.Bravo,
		strings.NewReader(
			"package example\n\ntype Foo struct{}\n",
		),
	)
	assertReport(t, "Bravo", false, nil, "", l)
}

func TestTypeCountUnexported(t *testing.T) {
	l := TypeCount(
		library.Charlie,
		strings.NewReader(
			"package example\n\ntype foo struct{}\n\ntype bar struct{}\n",
		),
	)
	assertReport(t, "Charlie", false, nil, "", l)
}
