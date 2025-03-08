package lint

import (
	"github.com/funtimecoding/go-library/pkg/lint/concern"
	stringLibrary "github.com/funtimecoding/go-library/pkg/strings"
	"strings"
	"testing"
)

func TestGo(t *testing.T) {
	l := Go(
		stringLibrary.Alfa,
		strings.NewReader(
			"package example\n\nimport (\n\t\"example.org/example/fmt\"\n)\n\nfunc Example() {\n\tfmt.Println(\"Hello friend.\")\n}\n",
		),
	)
	assertReport(
		t,
		"Alfa",
		true,
		[]*concern.Concern{
			{
				Key:      "single_multi_line_import",
				Text:     "single multi-line import",
				Line:     3,
				LineText: "import (",
			},
		},
		"package example\n\nimport \"example.org/example/fmt\"\n\nfunc Example() {\n\tfmt.Println(\"Hello friend.\")\n}\n",
		l,
	)
}
