package lint

import (
	"github.com/funtimecoding/go-library/pkg/lint/concern"
	stringLibrary "github.com/funtimecoding/go-library/pkg/strings"
	"strings"
	"testing"
)

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
			},
		},
		"---\nmyKey: myValue\n",
		l,
	)
}
