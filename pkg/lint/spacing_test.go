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
			Fixed: true,
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
			Fixed: true,
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
			Fixed: true,
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
			Fixed: true,
			},
			{
				Key:      lintConstant.BlankInsideFunctionKey,
				Text:     lintConstant.BlankInsideFunctionText,
				Path:     "Delta",
				Line:     5,
				LineText: "",
			Fixed: true,
			},
		},
		"package example\n\nfunc Example() {\n\tx := 1\n\ty := 2\n}\n",
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
	assertReport(
		t,
		"Kilo",
		false,
		nil,
		"package example\n\nfunc Example(x int) {\n\tswitch x {\n\tcase 1:\n\t\tif x > 0 {\n\t\t\tfmt.Println(\"a\")\n\t\t}\n\tdefault:\n\t\tfmt.Println(\"b\")\n\t}\n}\n",
		l,
	)
}

func TestSpacingClosingBraceBeforeComma(t *testing.T) {
	l := Spacing(
		library.Lima,
		strings.NewReader(
			"package example\n\nfunc Example(run func()) {\n\trun()\n}\n\nfunc Call() {\n\tExample(func() {\n\t\tif true {\n\t\t\tfmt.Println(\"a\")\n\t\t}\n\t})\n}\n",
		),
	)
	assertReport(
		t,
		"Lima",
		false,
		nil,
		"package example\n\nfunc Example(run func()) {\n\trun()\n}\n\nfunc Call() {\n\tExample(func() {\n\t\tif true {\n\t\t\tfmt.Println(\"a\")\n\t\t}\n\t})\n}\n",
		l,
	)
}

func TestSpacingContinueAsIdentifierPrefix(t *testing.T) {
	l := Spacing(
		library.Mike,
		strings.NewReader(
			"package example\n\nfunc Example(\n\tcontinueOnError bool,\n) {\n\tfmt.Println(continueOnError)\n}\n",
		),
	)
	assertReport(
		t,
		"Mike",
		false,
		nil,
		"package example\n\nfunc Example(\n\tcontinueOnError bool,\n) {\n\tfmt.Println(continueOnError)\n}\n",
		l,
	)
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

func TestSpacingBlankBetweenAssignments(t *testing.T) {
	l := Spacing(
		library.Quebec,
		strings.NewReader(
			"package example\n\nfunc Example() {\n\tx := 1\n\n\ty := 2\n}\n",
		),
	)
	assertReport(
		t,
		"Quebec",
		true,
		[]*concern.Concern{
			{
				Key:      lintConstant.BlankInsideFunctionKey,
				Text:     lintConstant.BlankInsideFunctionText,
				Path:     "Quebec",
				Line:     5,
				LineText: "",
			Fixed: true,
			},
		},
		"package example\n\nfunc Example() {\n\tx := 1\n\ty := 2\n}\n",
		l,
	)
}

func TestSpacingBlankAtStartOfFunction(t *testing.T) {
	l := Spacing(
		library.Romeo,
		strings.NewReader(
			"package example\n\nfunc Example() {\n\n\tx := 1\n}\n",
		),
	)
	assertReport(
		t,
		"Romeo",
		true,
		[]*concern.Concern{
			{
				Key:      lintConstant.BlankInsideFunctionKey,
				Text:     lintConstant.BlankInsideFunctionText,
				Path:     "Romeo",
				Line:     4,
				LineText: "",
			Fixed: true,
			},
		},
		"package example\n\nfunc Example() {\n\tx := 1\n}\n",
		l,
	)
}

func TestSpacingBlankAtEndOfFunction(t *testing.T) {
	l := Spacing(
		library.Sierra,
		strings.NewReader(
			"package example\n\nfunc Example() {\n\tx := 1\n\n}\n",
		),
	)
	assertReport(
		t,
		"Sierra",
		true,
		[]*concern.Concern{
			{
				Key:      lintConstant.BlankInsideFunctionKey,
				Text:     lintConstant.BlankInsideFunctionText,
				Path:     "Sierra",
				Line:     5,
				LineText: "",
			Fixed: true,
			},
		},
		"package example\n\nfunc Example() {\n\tx := 1\n}\n",
		l,
	)
}

func TestSpacingBlankBeforeReturnFirstInBlock(t *testing.T) {
	l := Spacing(
		library.Tango,
		strings.NewReader(
			"package example\n\nfunc Example() int {\n\n\treturn 1\n}\n",
		),
	)
	assertReport(
		t,
		"Tango",
		true,
		[]*concern.Concern{
			{
				Key:      lintConstant.BlankInsideFunctionKey,
				Text:     lintConstant.BlankInsideFunctionText,
				Path:     "Tango",
				Line:     4,
				LineText: "",
			Fixed: true,
			},
		},
		"package example\n\nfunc Example() int {\n\treturn 1\n}\n",
		l,
	)
}

func TestSpacingBlankBetweenAssignmentsInNestedBlock(t *testing.T) {
	l := Spacing(
		library.Uniform,
		strings.NewReader(
			"package example\n\nfunc Example() {\n\tif true {\n\t\tx := 1\n\n\t\ty := 2\n\t}\n}\n",
		),
	)
	assertReport(
		t,
		"Uniform",
		true,
		[]*concern.Concern{
			{
				Key:      lintConstant.BlankInsideFunctionKey,
				Text:     lintConstant.BlankInsideFunctionText,
				Path:     "Uniform",
				Line:     6,
				LineText: "",
			Fixed: true,
			},
		},
		"package example\n\nfunc Example() {\n\tif true {\n\t\tx := 1\n\t\ty := 2\n\t}\n}\n",
		l,
	)
}

func TestSpacingBlankAfterControlBeforeStatement(t *testing.T) {
	l := Spacing(
		library.Victor,
		strings.NewReader(
			"package example\n\nfunc Example() {\n\tif true {\n\t\tfmt.Println(\"a\")\n\t}\n\n\tfmt.Println(\"b\")\n}\n",
		),
	)
	assertReport(
		t,
		"Victor",
		false,
		nil,
		"package example\n\nfunc Example() {\n\tif true {\n\t\tfmt.Println(\"a\")\n\t}\n\n\tfmt.Println(\"b\")\n}\n",
		l,
	)
}

func TestSpacingBlankBeforeIfInFunction(t *testing.T) {
	l := Spacing(
		library.Whiskey,
		strings.NewReader(
			"package example\n\nfunc Example() {\n\tx := 1\n\n\tif x > 0 {\n\t\tfmt.Println(x)\n\t}\n}\n",
		),
	)
	assertReport(
		t,
		"Whiskey",
		false,
		nil,
		"package example\n\nfunc Example() {\n\tx := 1\n\n\tif x > 0 {\n\t\tfmt.Println(x)\n\t}\n}\n",
		l,
	)
}

func TestSpacingBlankInFuncLiteral(t *testing.T) {
	l := Spacing(
		library.Xray,
		strings.NewReader(
			"package example\n\nfunc Example() {\n\tf := func() {\n\t\tx := 1\n\n\t\ty := 2\n\t}\n\tf()\n}\n",
		),
	)
	assertReport(
		t,
		"Xray",
		true,
		[]*concern.Concern{
			{
				Key:      lintConstant.BlankInsideFunctionKey,
				Text:     lintConstant.BlankInsideFunctionText,
				Path:     "Xray",
				Line:     6,
				LineText: "",
			Fixed: true,
			},
		},
		"package example\n\nfunc Example() {\n\tf := func() {\n\t\tx := 1\n\t\ty := 2\n\t}\n\tf()\n}\n",
		l,
	)
}

func TestSpacingCompositeLiteralCloserDoesNotLeakDepth(t *testing.T) {
	l := Spacing(
		"composite",
		strings.NewReader(
			"package example\n\nfunc First() {\n\t_ = []string{\n\t\t\"a\",\n\t}\n}\n\nfunc Second() {}\n",
		),
	)
	assertReport(
		t,
		"composite",
		false,
		nil,
		"package example\n\nfunc First() {\n\t_ = []string{\n\t\t\"a\",\n\t}\n}\n\nfunc Second() {}\n",
		l,
	)
}

func TestSpacingBlankBeforeCommentBeforeControl(t *testing.T) {
	l := Spacing(
		library.Zulu,
		strings.NewReader(
			"package example\n\nfunc Example() {\n\tx := 1\n\n\t// check\n\tif x > 0 {\n\t\tfmt.Println(x)\n\t}\n}\n",
		),
	)
	assertReport(
		t,
		"Zulu",
		false,
		nil,
		"package example\n\nfunc Example() {\n\tx := 1\n\n\t// check\n\tif x > 0 {\n\t\tfmt.Println(x)\n\t}\n}\n",
		l,
	)
}

func TestSpacingBlankBeforeCommentBeforeAssignment(t *testing.T) {
	l := Spacing(
		library.Yankee,
		strings.NewReader(
			"package example\n\nfunc Example() {\n\tx := 1\n\n\t// note\n\ty := 2\n}\n",
		),
	)
	assertReport(
		t,
		"Yankee",
		false,
		nil,
		"package example\n\nfunc Example() {\n\tx := 1\n\n\t// note\n\ty := 2\n}\n",
		l,
	)
}

func TestSpacingMultiLineConditionBlankAfterControl(t *testing.T) {
	l := Spacing(
		"multiline",
		strings.NewReader(
			"package example\n\nfunc Example() {\n\tif strings.Contains(\n\t\ts,\n\t\t\"x\",\n\t) {\n\t\tcontinue\n\t}\n\n\tx := 1\n\t_ = x\n}\n",
		),
	)
	assertReport(
		t,
		"multiline",
		false,
		nil,
		"package example\n\nfunc Example() {\n\tif strings.Contains(\n\t\ts,\n\t\t\"x\",\n\t) {\n\t\tcontinue\n\t}\n\n\tx := 1\n\t_ = x\n}\n",
		l,
	)
}

func TestSpacingBlankBeforeDefer(t *testing.T) {
	l := Spacing(
		"defer",
		strings.NewReader(
			"package example\n\nfunc Example() {\n\tx, e := f()\n\n\tdefer g()\n\t_ = x\n\t_ = e\n}\n",
		),
	)
	assertReport(
		t,
		"defer",
		false,
		nil,
		"package example\n\nfunc Example() {\n\tx, e := f()\n\n\tdefer g()\n\t_ = x\n\t_ = e\n}\n",
		l,
	)
}

func TestSpacingMissingBlankBetweenFunctions(t *testing.T) {
	l := Spacing(
		"decl",
		strings.NewReader(
			"package example\n\nfunc First() {}\nfunc Second() {}\n",
		),
	)
	assertReport(
		t,
		"decl",
		true,
		[]*concern.Concern{
			{
				Key:      lintConstant.MissingBlankBeforeDeclarationKey,
				Text:     lintConstant.MissingBlankBeforeDeclarationText,
				Path:     "decl",
				Line:     4,
				LineText: "func Second() {}",
			Fixed: true,
			},
		},
		"package example\n\nfunc First() {}\n\nfunc Second() {}\n",
		l,
	)
}

func TestSpacingConsecutiveTopLevelVarNoBlankAdded(t *testing.T) {
	l := Spacing(
		"topvar",
		strings.NewReader(
			"package example\n\nvar A = 1\nvar B = 2\nvar C = 3\n",
		),
	)
	assertReport(
		t,
		"topvar",
		false,
		nil,
		"package example\n\nvar A = 1\nvar B = 2\nvar C = 3\n",
		l,
	)
}

func TestSpacingSpuriousBlankBetweenTopLevelVarsRemoved(t *testing.T) {
	l := Spacing(
		"topvar2",
		strings.NewReader(
			"package example\n\nvar A = 1\n\nvar B = 2\n\nvar C = 3\n",
		),
	)
	assertReport(
		t,
		"topvar2",
		true,
		[]*concern.Concern{
			{
				Key:      lintConstant.ExtraneousTopLevelBlankKey,
				Text:     lintConstant.ExtraneousTopLevelBlankText,
				Path:     "topvar2",
				Line:     4,
				LineText: "",
			Fixed: true,
			},
			{
				Key:      lintConstant.ExtraneousTopLevelBlankKey,
				Text:     lintConstant.ExtraneousTopLevelBlankText,
				Path:     "topvar2",
				Line:     6,
				LineText: "",
			Fixed: true,
			},
		},
		"package example\n\nvar A = 1\nvar B = 2\nvar C = 3\n",
		l,
	)
}

func TestSpacingBlankBeforeCommentAfterControlIdempotent(t *testing.T) {
	input := "package example\n\nfunc Example() {\n\tif true {\n\t\tx := 1\n\t\t_ = x\n\t}\n\n\t// note\n\ty := 2\n\t_ = y\n}\n"
	l := Spacing("idempotent", strings.NewReader(input))
	assertReport(t, "idempotent", false, nil, input, l)
}

func TestSpacingBlankAfterControlBeforeRegularStatementPreserved(t *testing.T) {
	l := Spacing(
		"ctrlblank",
		strings.NewReader(
			"package example\n\nfunc Example() {\n\tif true {\n\t\tx := 1\n\t\t_ = x\n\t}\n\n\ty := 2\n\t_ = y\n}\n",
		),
	)
	assertReport(
		t,
		"ctrlblank",
		false,
		nil,
		"package example\n\nfunc Example() {\n\tif true {\n\t\tx := 1\n\t\t_ = x\n\t}\n\n\ty := 2\n\t_ = y\n}\n",
		l,
	)
}

func TestSpacingBlankAfterNestedControlBeforeRegularStatementPreserved(t *testing.T) {
	l := Spacing(
		"nestedctrl",
		strings.NewReader(
			"package example\n\nfunc Example() {\n\tfor _, s := range []string{\n\t\t\"a\",\n\t\t\"b\",\n\t} {\n\t\tif !contains(s) {\n\t\t\tadd(s)\n\t\t}\n\t}\n\n\tif !contains(\"c\") {\n\t\tadd(\"c\")\n\t}\n\n\tresult := count()\n\t_ = result\n}\n",
		),
	)
	assertReport(
		t,
		"nestedctrl",
		false,
		nil,
		"package example\n\nfunc Example() {\n\tfor _, s := range []string{\n\t\t\"a\",\n\t\t\"b\",\n\t} {\n\t\tif !contains(s) {\n\t\t\tadd(s)\n\t\t}\n\t}\n\n\tif !contains(\"c\") {\n\t\tadd(\"c\")\n\t}\n\n\tresult := count()\n\t_ = result\n}\n",
		l,
	)
}

func TestSpacingBlankAfterCommentRemoved(t *testing.T) {
	l := Spacing(
		"aftercomment",
		strings.NewReader(
			"package example\n\nfunc Example() {\n\t// note\n\n\tif true {\n\t\treturn\n\t}\n}\n",
		),
	)
	assertReport(
		t,
		"aftercomment",
		true,
		[]*concern.Concern{
			{
				Key:      lintConstant.BlankInsideFunctionKey,
				Text:     lintConstant.BlankInsideFunctionText,
				Path:     "aftercomment",
				Line:     5,
				LineText: "",
			Fixed: true,
			},
		},
		"package example\n\nfunc Example() {\n\t// note\n\tif true {\n\t\treturn\n\t}\n}\n",
		l,
	)
}

func TestSpacingInterFunctionBlankAfterNestedFuncLiteral(t *testing.T) {
	l := Spacing(
		"interfunc",
		strings.NewReader(
			"package example\n\nfunc First() {\n\tn(\n\t\tfunc(m *M) {\n\t\t\tm.H(\n\t\t\t\t\"/\",\n\t\t\t\tfunc(\n\t\t\t\t\tw W,\n\t\t\t\t\t_ R,\n\t\t\t\t) {\n\t\t\t\t\tw.Write()\n\t\t\t\t},\n\t\t\t)\n\t\t},\n\t)\n}\n\nfunc Second() {}\n",
		),
	)
	assertReport(
		t,
		"interfunc",
		false,
		nil,
		"package example\n\nfunc First() {\n\tn(\n\t\tfunc(m *M) {\n\t\t\tm.H(\n\t\t\t\t\"/\",\n\t\t\t\tfunc(\n\t\t\t\t\tw W,\n\t\t\t\t\t_ R,\n\t\t\t\t) {\n\t\t\t\t\tw.Write()\n\t\t\t\t},\n\t\t\t)\n\t\t},\n\t)\n}\n\nfunc Second() {}\n",
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
			Fixed: true,
			},
		},
		"package example\n\nfunc Example() {\n\tif true {\n\t\tfmt.Println(\"a\")\n\t}\n\n\tfmt.Println(\"b\")\n}\n",
		l,
	)
}
