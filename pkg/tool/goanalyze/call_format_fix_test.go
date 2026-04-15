package goanalyze

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"path/filepath"
	"testing"
)

func TestCallFormatFix(t *testing.T) {
	directory := writeCallFormatTestModule(t)
	runCallFormatFixInDirectory(directory)
	t.Run(
		"LongSingleLine",
		func(t *testing.T) {
			assert.String(
				t,
				"package example\n\nfunc LongSingleLine() {\n\ttwoArgs(\n\t\t\"something-long-enough\",\n\t\t\"to-push-this-well-past-the-eighty-character-column-limit\",\n\t)\n}\n\nfunc twoArgs(a, b string) {}\n",
				readFile(
					t,
					filepath.Join(
						directory,
						"long_single_line.go",
					),
				),
			)
		},
	)
	t.Run(
		"SharedLine",
		func(t *testing.T) {
			assert.String(
				t,
				"package example\n\ntype Options struct {\n\tValue int\n}\n\nfunc SharedLine() {\n\twithStruct(\n\t\t\"name\",\n\t\tOptions{\n\t\t\tValue: 1,\n\t\t},\n\t)\n}\n\nfunc withStruct(a string, b Options) {}\n",
				readFile(
					t,
					filepath.Join(directory, "shared_line.go"),
				),
			)
		},
	)
	t.Run(
		"FirstArgOnParenLine",
		func(t *testing.T) {
			assert.String(
				t,
				"package example\n\nfunc FirstArgOnParenLine() {\n\twithMap(\n\t\t\"name\",\n\t\tmap[string]any{\n\t\t\t\"key\": \"value\",\n\t\t},\n\t)\n}\n\nfunc withMap(a string, b map[string]any) {}\n",
				readFile(
					t,
					filepath.Join(
						directory,
						"first_arg_on_paren_line.go",
					),
				),
			)
		},
	)
	t.Run(
		"NestedIndent",
		func(t *testing.T) {
			assert.String(
				t,
				"package example\n\nfunc NestedIndent() {\n\tif true {\n\t\tif true {\n\t\t\ttwoArgs(\n\t\t\t\t\"something-long-enough-to-exceed\",\n\t\t\t\t\"the-eighty-character-limit-at-this-indent-level\",\n\t\t\t)\n\t\t}\n\t}\n}\n",
				readFile(
					t,
					filepath.Join(directory, "nested_indent.go"),
				),
			)
		},
	)
	t.Run(
		"CompliantUntouched",
		func(t *testing.T) {
			assert.String(
				t,
				"package example\n\nfunc Compliant() {\n\ttwoArgs(\"alpha\", \"bravo\")\n\ttwoArgs(\n\t\t\"alpha\",\n\t\t\"bravo\",\n\t)\n}\n",
				readFile(
					t,
					filepath.Join(directory, "compliant.go"),
				),
			)
		},
	)
	t.Run(
		"DeepMethodCall",
		func(t *testing.T) {
			assert.String(
				t,
				"package example\n\ntype Logger struct{}\n\nfunc (l *Logger) Structured(args ...string) {}\n\ntype Poller struct {\n\tlogger *Logger\n}\n\nfunc (p *Poller) Run() {\n\tdefer func() {\n\t\tif v := recover(); v != nil {\n\t\t\tp.logger.Structured(\n\t\t\t\t\"recover failed\",\n\t\t\t\t\"error\",\n\t\t\t\t\"value\",\n\t\t\t)\n\t\t}\n\t}()\n}\n",
				readFile(
					t,
					filepath.Join(directory, "deep_method.go"),
				),
			)
		},
	)
	t.Run(
		"MultipleViolationsInOneCall",
		func(t *testing.T) {
			assert.String(
				t,
				"package example\n\nfunc MultipleViolations() {\n\tfourArgs(\n\t\t\"adopted\",\n\t\t\"bravo\",\n\t\t\"charlie\",\n\t\t\"delta\",\n\t)\n}\n\nfunc fourArgs(a, b, c, d string) {}\n",
				readFile(
					t,
					filepath.Join(
						directory,
						"multiple_violations.go",
					),
				),
			)
		},
	)
}

func writeCallFormatTestModule(t *testing.T) string {
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
		"long_single_line.go",
		"package example\n\nfunc LongSingleLine() {\n\ttwoArgs(\"something-long-enough\", \"to-push-this-well-past-the-eighty-character-column-limit\")\n}\n\nfunc twoArgs(a, b string) {}\n",
	)
	writeFile(
		t,
		directory,
		"shared_line.go",
		"package example\n\ntype Options struct {\n\tValue int\n}\n\nfunc SharedLine() {\n\twithStruct(\n\t\t\"name\", Options{\n\t\t\tValue: 1,\n\t\t},\n\t)\n}\n\nfunc withStruct(a string, b Options) {}\n",
	)
	writeFile(
		t,
		directory,
		"first_arg_on_paren_line.go",
		"package example\n\nfunc FirstArgOnParenLine() {\n\twithMap(\"name\",\n\t\tmap[string]any{\n\t\t\t\"key\": \"value\",\n\t\t},\n\t)\n}\n\nfunc withMap(a string, b map[string]any) {}\n",
	)
	writeFile(
		t,
		directory,
		"nested_indent.go",
		"package example\n\nfunc NestedIndent() {\n\tif true {\n\t\tif true {\n\t\t\ttwoArgs(\"something-long-enough-to-exceed\", \"the-eighty-character-limit-at-this-indent-level\")\n\t\t}\n\t}\n}\n",
	)
	writeFile(
		t,
		directory,
		"compliant.go",
		"package example\n\nfunc Compliant() {\n\ttwoArgs(\"alpha\", \"bravo\")\n\ttwoArgs(\n\t\t\"alpha\",\n\t\t\"bravo\",\n\t)\n}\n",
	)
	writeFile(
		t,
		directory,
		"multiple_violations.go",
		"package example\n\nfunc MultipleViolations() {\n\tfourArgs(\n\t\t\"adopted\", \"bravo\",\n\t\t\"charlie\", \"delta\",\n\t)\n}\n\nfunc fourArgs(a, b, c, d string) {}\n",
	)
	writeFile(
		t,
		directory,
		"deep_method.go",
		"package example\n\ntype Logger struct{}\n\nfunc (l *Logger) Structured(args ...string) {}\n\ntype Poller struct {\n\tlogger *Logger\n}\n\nfunc (p *Poller) Run() {\n\tdefer func() {\n\t\tif v := recover(); v != nil {\n\t\t\tp.logger.Structured(\n\t\t\t\t\"recover failed\",\n\t\t\t\t\"error\", \"value\",\n\t\t\t)\n\t\t}\n\t}()\n}\n",
	)

	return directory
}

func runCallFormatFixInDirectory(directory string) {
	runCallFormatFixWithDirectory([]string{"./..."}, directory)
}
