package lint

import (
	"github.com/funtimecoding/go-library/pkg/lint/concern"
	stringLibrary "github.com/funtimecoding/go-library/pkg/strings"
	"strings"
	"testing"
)

func TestSpacingBlankBeforeControl(t *testing.T) {
	l := Spacing(
		stringLibrary.Alfa,
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
				Key:      "missing_blank_before_control",
				Text:     "Missing blank line before control block",
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
		stringLibrary.Bravo,
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
				Key:      "missing_blank_before_exit",
				Text:     "Missing blank line before return/break/continue",
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
		stringLibrary.Charlie,
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
				Key:      "missing_blank_after_control",
				Text:     "Missing blank line after control block",
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
		stringLibrary.Echo,
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
		stringLibrary.Foxtrot,
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
		stringLibrary.Golf,
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
		stringLibrary.Hotel,
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
		stringLibrary.Delta,
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
				Key:      "extraneous_blank_line",
				Text:     "Extraneous blank line",
				Path:     "Delta",
				Line:     6,
				LineText: "",
			},
		},
		"package example\n\nfunc Example() {\n\tx := 1\n\n\ty := 2\n}\n",
		l,
	)
}
