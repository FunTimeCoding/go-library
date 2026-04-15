package goanalyze

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestReplaceSegment(t *testing.T) {
	t.Run(
		"CamelCaseSingleLetter",
		func(t *testing.T) {
			assert.String(t, "fooR", replaceSegment("fooRef", "ref", "r"))
		},
	)
	t.Run(
		"CamelCaseMultiWord",
		func(t *testing.T) {
			assert.String(
				t,
				"fooReference",
				replaceSegment("fooRef", "ref", "reference"),
			)
		},
	)
	t.Run(
		"PascalCaseMultiWord",
		func(t *testing.T) {
			assert.String(
				t,
				"FooReference",
				replaceSegment("FooRef", "ref", "reference"),
			)
		},
	)
	t.Run(
		"SnakeCaseMultiWord",
		func(t *testing.T) {
			assert.String(
				t,
				"foo_reference",
				replaceSegment("foo_ref", "ref", "reference"),
			)
		},
	)
	t.Run(
		"SnakeCaseUnderscoreExpansion",
		func(t *testing.T) {
			assert.String(
				t,
				"foo_model_context",
				replaceSegment("foo_mcp", "mcp", "model_context"),
			)
		},
	)
	t.Run(
		"CamelCaseUnderscoreExpansion",
		func(t *testing.T) {
			assert.String(
				t,
				"fooModelContext",
				replaceSegment("fooMcp", "mcp", "model_context"),
			)
		},
	)
	t.Run(
		"SingleWordReplacement",
		func(t *testing.T) {
			assert.String(t, "r", replaceSegment("ref", "ref", "r"))
		},
	)
	t.Run(
		"NoMatch",
		func(t *testing.T) {
			assert.String(
				t,
				"fooBar",
				replaceSegment("fooBar", "ref", "reference"),
			)
		},
	)
}
