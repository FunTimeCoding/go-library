package lint

import (
	"github.com/funtimecoding/go-library/pkg/lint/concern"
	"github.com/funtimecoding/go-library/pkg/strings/upper"
	"strings"
	"testing"
)

func TestMarkupClean(t *testing.T) {
	l := Markup(
		upper.Bravo,
		strings.NewReader("---\nmyKey: myValue\n"),
	)
	assertReport(t, "Bravo", false, nil, "", l)
}

func TestMarkup(t *testing.T) {
	l := Markup(
		upper.Alfa,
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
				Type:     concern.Line,
				Line:     1,
				LineText: "myKey: myValue",
				Fixed:    true,
			},
		},
		"---\nmyKey: myValue\n",
		l,
	)
}
