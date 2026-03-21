package lint

import (
	"github.com/funtimecoding/go-library/pkg/lint/concern"
	lintConstant "github.com/funtimecoding/go-library/pkg/lint/constant"
	stringLibrary "github.com/funtimecoding/go-library/pkg/strings"
	"strings"
	"testing"
)

func TestForbiddenImportSingleLine(t *testing.T) {
	l := ForbiddenImport(
		stringLibrary.Alfa,
		strings.NewReader(
			"package example\n\nimport \"flag\"\n\nfunc Example() {\n\tflag.Parse()\n}\n",
		),
	)
	assertReport(
		t,
		"Alfa",
		true,
		[]*concern.Concern{
			{
				Key:      lintConstant.ForbiddenImportKey,
				Text:     lintConstant.ForbiddenImportText,
				Path:     "Alfa",
				Line:     3,
				LineText: `import "flag"`,
			},
		},
		"",
		l,
	)
}

func TestForbiddenImportBlock(t *testing.T) {
	l := ForbiddenImport(
		stringLibrary.Bravo,
		strings.NewReader(
			"package example\n\nimport (\n\t\"flag\"\n\t\"fmt\"\n)\n\nfunc Example() {\n\tflag.Parse()\n\tfmt.Println(\"ok\")\n}\n",
		),
	)
	assertReport(
		t,
		"Bravo",
		true,
		[]*concern.Concern{
			{
				Key:      lintConstant.ForbiddenImportKey,
				Text:     lintConstant.ForbiddenImportText,
				Path:     "Bravo",
				Line:     4,
				LineText: "\t\"flag\"",
			},
		},
		"",
		l,
	)
}

func TestForbiddenImportClean(t *testing.T) {
	l := ForbiddenImport(
		stringLibrary.Charlie,
		strings.NewReader(
			"package example\n\nimport (\n\t\"fmt\"\n\t\"github.com/spf13/pflag\"\n)\n\nfunc Example() {\n\tfmt.Println(pflag.NFlag())\n}\n",
		),
	)
	assertReport(t, "Charlie", false, nil, "", l)
}
