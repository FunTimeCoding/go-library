package lint

import (
	"github.com/funtimecoding/go-library/pkg/lint/concern"
	lintConstant "github.com/funtimecoding/go-library/pkg/lint/constant"
	library "github.com/funtimecoding/go-library/pkg/strings"
	"strings"
	"testing"
)

func TestSpacingBlankBeforeControl(t *testing.T) {
	l := Spacing(
		library.Alfa,
		strings.NewReader(
			"package example\n\nfunc Example() {\n\tx := 1\n\tif x > 0 {\n\t\tfmt.Println(x)\n\t}\n}\n",
		),
	)
	assertReport(
		t,
		"Alfa",
		true,
		[]*concern.Concern{
			{
				Key:      lintConstant.MissingBlankBeforeControlKey,
				Text:     lintConstant.MissingBlankBeforeControlText,
				Path:     "Alfa",
				Line:     5,
				LineText: "\tif x > 0 {",
			},
		},
		"package example\n\nfunc Example() {\n\tx := 1\n\n\tif x > 0 {\n\t\tfmt.Println(x)\n\t}\n}\n",
		l,
	)
}

func TestSpacingBlankBeforeReturn(t *testing.T) {
	l := Spacing(
		library.Bravo,
		strings.NewReader(
			"package example\n\nfunc Example() int {\n\tx := 1\n\treturn x\n}\n",
		),
	)
	assertReport(
		t,
		"Bravo",
		true,
		[]*concern.Concern{
			{
				Key:      lintConstant.MissingBlankBeforeExitKey,
				Text:     lintConstant.MissingBlankBeforeExitText,
				Path:     "Bravo",
				Line:     5,
				LineText: "\treturn x",
			},
		},
		"package example\n\nfunc Example() int {\n\tx := 1\n\n\treturn x\n}\n",
		l,
	)
}

func TestSpacingBlankAfterControl(t *testing.T) {
	l := Spacing(
		library.Charlie,
		strings.NewReader(
			"package example\n\nfunc Example() {\n\tx := 1\n\n\tif true {\n\t\tfmt.Println(\"a\")\n\t}\n\tfmt.Println(\"b\")\n}\n",
		),
	)
	assertReport(
		t,
		"Charlie",
		true,
		[]*concern.Concern{
			{
				Key:      lintConstant.MissingBlankAfterControlKey,
				Text:     lintConstant.MissingBlankAfterControlText,
				Path:     "Charlie",
				Line:     9,
				LineText: "\tfmt.Println(\"b\")",
			},
		},
		"package example\n\nfunc Example() {\n\tx := 1\n\n\tif true {\n\t\tfmt.Println(\"a\")\n\t}\n\n\tfmt.Println(\"b\")\n}\n",
		l,
	)
}

func TestSpacingReturnFirstInBlock(t *testing.T) {
	l := Spacing(
		library.Echo,
		strings.NewReader(
			"package example\n\nfunc Example() int {\n\treturn 1\n}\n",
		),
	)
	assertReport(
		t,
		"Echo",
		false,
		nil,
		"package example\n\nfunc Example() int {\n\treturn 1\n}\n",
		l,
	)
}

func TestSpacingControlFirstInBlock(t *testing.T) {
	l := Spacing(
		library.Foxtrot,
		strings.NewReader(
			"package example\n\nfunc Example() {\n\tif true {\n\t\tfmt.Println(\"a\")\n\t}\n}\n",
		),
	)
	assertReport(
		t,
		"Foxtrot",
		false,
		nil,
		"package example\n\nfunc Example() {\n\tif true {\n\t\tfmt.Println(\"a\")\n\t}\n}\n",
		l,
	)
}

func TestSpacingConsecutiveClosingBraces(t *testing.T) {
	l := Spacing(
		library.Golf,
		strings.NewReader(
			"package example\n\nfunc Example() {\n\tif true {\n\t\tif true {\n\t\t\tfmt.Println(\"a\")\n\t\t}\n\t}\n}\n",
		),
	)
	assertReport(
		t,
		"Golf",
		false,
		nil,
		"package example\n\nfunc Example() {\n\tif true {\n\t\tif true {\n\t\t\tfmt.Println(\"a\")\n\t\t}\n\t}\n}\n",
		l,
	)
}

func TestSpacingElse(t *testing.T) {
	l := Spacing(
		library.Hotel,
		strings.NewReader(
			"package example\n\nfunc Example() {\n\tif true {\n\t\tfmt.Println(\"a\")\n\t} else {\n\t\tfmt.Println(\"b\")\n\t}\n}\n",
		),
	)
	assertReport(
		t,
		"Hotel",
		false,
		nil,
		"package example\n\nfunc Example() {\n\tif true {\n\t\tfmt.Println(\"a\")\n\t} else {\n\t\tfmt.Println(\"b\")\n\t}\n}\n",
		l,
	)
}

func TestSpacingExtraneousBlanks(t *testing.T) {
	l := Spacing(
		library.Delta,
		strings.NewReader(
			"package example\n\nfunc Example() {\n\tx := 1\n\n\n\ty := 2\n}\n",
		),
	)
	assertReport(
		t,
		"Delta",
		true,
		[]*concern.Concern{
			{
				Key:      lintConstant.ExtraneousBlankLineKey,
				Text:     lintConstant.ExtraneousBlankLineText,
				Path:     "Delta",
				Line:     6,
				LineText: "",
			},
		},
		"package example\n\nfunc Example() {\n\tx := 1\n\n\ty := 2\n}\n",
		l,
	)
}

func TestSpacingClean(t *testing.T) {
	l := Spacing(
		library.Juliett,
		strings.NewReader(
			"package example\n\nfunc Example() {\n\tfmt.Println(\"a\")\n}\n",
		),
	)
	assertReport(
		t,
		"Juliett",
		false,
		nil,
		"package example\n\nfunc Example() {\n\tfmt.Println(\"a\")\n}\n",
		l,
	)
}

func TestSpacingClosingBraceBeforeDefault(t *testing.T) {
	l := Spacing(
		library.Kilo,
		strings.NewReader(
			"package example\n\nfunc Example(x int) {\n\tswitch x {\n\tcase 1:\n\t\tif x > 0 {\n\t\t\tfmt.Println(\"a\")\n\t\t}\n\tdefault:\n\t\tfmt.Println(\"b\")\n\t}\n}\n",
		),
	)
	assertReport(t, "Kilo", false, nil, "package example\n\nfunc Example(x int) {\n\tswitch x {\n\tcase 1:\n\t\tif x > 0 {\n\t\t\tfmt.Println(\"a\")\n\t\t}\n\tdefault:\n\t\tfmt.Println(\"b\")\n\t}\n}\n", l)
}

func TestSpacingClosingBraceBeforeComma(t *testing.T) {
	l := Spacing(
		library.Lima,
		strings.NewReader(
			"package example\n\nfunc Example(run func()) {\n\trun()\n}\n\nfunc Call() {\n\tExample(func() {\n\t\tif true {\n\t\t\tfmt.Println(\"a\")\n\t\t}\n\t})\n}\n",
		),
	)
	assertReport(t, "Lima", false, nil, "package example\n\nfunc Example(run func()) {\n\trun()\n}\n\nfunc Call() {\n\tExample(func() {\n\t\tif true {\n\t\t\tfmt.Println(\"a\")\n\t\t}\n\t})\n}\n", l)
}

func TestSpacingCommentBeforeControl(t *testing.T) {
	l := Spacing(
		library.Mike,
		strings.NewReader(
			"package example\n\nfunc Example() {\n\tx := 1\n\t// Check condition\n\tif x > 0 {\n\t\tfmt.Println(x)\n\t}\n}\n",
		),
	)
	assertReport(
		t,
		"Mike",
		false,
		nil,
		"package example\n\nfunc Example() {\n\tx := 1\n\t// Check condition\n\tif x > 0 {\n\t\tfmt.Println(x)\n\t}\n}\n",
		l,
	)
}

func TestSpacingCompositeLiteralInFunc(t *testing.T) {
	l := Spacing(
		library.November,
		strings.NewReader(
			"package example\n\nfunc Example() {\n\tx := map[string]int{\n\t\t\"a\": 1,\n\t}\n\ty := 2\n}\n",
		),
	)
	assertReport(
		t,
		"November",
		false,
		nil,
		"package example\n\nfunc Example() {\n\tx := map[string]int{\n\t\t\"a\": 1,\n\t}\n\ty := 2\n}\n",
		l,
	)
}

func TestSpacingRawStringLiteral(t *testing.T) {
	l := Spacing(
		library.Oscar,
		strings.NewReader(
			"package example\n\nvar Template = `\nif (x) {\n\treturn false;\n}\ny = 1;\n`\n",
		),
	)
	assertReport(
		t,
		"Oscar",
		false,
		nil,
		"package example\n\nvar Template = `\nif (x) {\n\treturn false;\n}\ny = 1;\n`\n",
		l,
	)
}

func TestSpacingVarBlockCompositeLiteral(t *testing.T) {
	l := Spacing(
		library.Papa,
		strings.NewReader(
			"package example\n\nvar (\n\tStatuses = []string{\n\t\t\"open\",\n\t\t\"closed\",\n\t}\n\tOther = \"other\"\n)\n",
		),
	)
	assertReport(
		t,
		"Papa",
		false,
		nil,
		"package example\n\nvar (\n\tStatuses = []string{\n\t\t\"open\",\n\t\t\"closed\",\n\t}\n\tOther = \"other\"\n)\n",
		l,
	)
}

func TestSpacingDoubleBlankAfterClosingBrace(t *testing.T) {
	l := Spacing(
		library.India,
		strings.NewReader(
			"package example\n\nfunc Example() {\n\tif true {\n\t\tfmt.Println(\"a\")\n\t}\n\n\n\tfmt.Println(\"b\")\n}\n",
		),
	)
	assertReport(
		t,
		"India",
		true,
		[]*concern.Concern{
			{
				Key:      lintConstant.ExtraneousBlankLineKey,
				Text:     lintConstant.ExtraneousBlankLineText,
				Path:     "India",
				Line:     8,
				LineText: "",
			},
		},
		"package example\n\nfunc Example() {\n\tif true {\n\t\tfmt.Println(\"a\")\n\t}\n\n\tfmt.Println(\"b\")\n}\n",
		l,
	)
}
