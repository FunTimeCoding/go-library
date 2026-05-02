package gofix

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"path/filepath"
	"testing"
)

func TestVariableNamingFix(t *testing.T) {
	directory := writeVariableNamingTestModule(t)
	var r results
	runVariableNamingFixWithDirectory([]string{"./..."}, directory, false, &r)
	t.Run(
		"WrongSingleLetter",
		func(t *testing.T) {
			assert.String(
				t,
				"package example\n\nimport \"fmt\"\n\nfunc WrongSingle() {\n\te := fmt.Errorf(\"test\")\n\t_ = e\n}\n",
				readFile(
					t,
					filepath.Join(directory, "wrong_single.go"),
				),
			)
		},
	)
	t.Run(
		"ErrorRenamed",
		func(t *testing.T) {
			assert.String(
				t,
				"package example\n\nimport \"fmt\"\n\nfunc ErrorRenamed() {\n\te := fmt.Errorf(\"test\")\n\t_ = e\n}\n",
				readFile(
					t,
					filepath.Join(directory, "error_renamed.go"),
				),
			)
		},
	)
	t.Run(
		"ErrorChainRenamed",
		func(t *testing.T) {
			assert.String(
				t,
				"package example\n\nimport \"fmt\"\n\nfunc ErrorChainRenamed() {\n\te := fmt.Errorf(\"first\")\n\tf := fmt.Errorf(\"second\")\n\t_ = e\n\t_ = f\n}\n",
				readFile(
					t,
					filepath.Join(
						directory,
						"error_chain_renamed.go",
					),
				),
			)
		},
	)
	t.Run(
		"CorrectUntouched",
		func(t *testing.T) {
			assert.String(
				t,
				"package example\n\nimport \"fmt\"\n\nfunc CorrectUntouched() {\n\te := fmt.Errorf(\"test\")\n\ts := \"hello\"\n\t_ = e\n\t_ = s\n}\n",
				readFile(
					t,
					filepath.Join(directory, "correct.go"),
				),
			)
		},
	)
}

func writeVariableNamingTestModule(t *testing.T) string {
	t.Helper()
	directory := t.TempDir()
	writeFile(
		t,
		directory,
		"go.mod",
		"module example\n\ngo 1.22\n",
	)
	writeFile(
		t,
		directory,
		"wrong_single.go",
		"package example\n\nimport \"fmt\"\n\nfunc WrongSingle() {\n\tx := fmt.Errorf(\"test\")\n\t_ = x\n}\n",
	)
	writeFile(
		t,
		directory,
		"error_renamed.go",
		"package example\n\nimport \"fmt\"\n\nfunc ErrorRenamed() {\n\terr := fmt.Errorf(\"test\")\n\t_ = err\n}\n",
	)
	writeFile(
		t,
		directory,
		"error_chain_renamed.go",
		"package example\n\nimport \"fmt\"\n\nfunc ErrorChainRenamed() {\n\terr := fmt.Errorf(\"first\")\n\terr2 := fmt.Errorf(\"second\")\n\t_ = err\n\t_ = err2\n}\n",
	)
	writeFile(
		t,
		directory,
		"correct.go",
		"package example\n\nimport \"fmt\"\n\nfunc CorrectUntouched() {\n\te := fmt.Errorf(\"test\")\n\ts := \"hello\"\n\t_ = e\n\t_ = s\n}\n",
	)

	return directory
}
