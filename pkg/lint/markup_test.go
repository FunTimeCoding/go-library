package lint

import (
	"github.com/funtimecoding/go-library/pkg/lint/concern"
	stringLibrary "github.com/funtimecoding/go-library/pkg/strings"
	"strings"
	"testing"
)

func TestMarkupClean(t *testing.T) {
	l := Markup(
		stringLibrary.Bravo,
		strings.NewReader("---\nmyKey: myValue\n"),
	)
	assertReport(t, "Bravo", false, nil, "", l)
}

func TestMarkup(t *testing.T) {
	l := Markup(
		stringLibrary.Alfa,
		strings.NewReader("myKey: myValue\n"),
	)
	assertReport(
		t,
		"Alfa",
		true,
		[]*concern.Concern{
			{
				Key:      "front_matter_delimiter",
				Text:     "No front matter delimiter",
				Path:     "Alfa",
				Line:     1,
				LineText: "myKey: myValue",
				Fixed:    true,
			},
		},
		"---\nmyKey: myValue\n",
		l,
	)
}
