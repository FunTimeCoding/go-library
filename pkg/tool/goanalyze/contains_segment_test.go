package goanalyze

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestContainsSegment(t *testing.T) {
	t.Run(
		"SingleLetterMatch",
		func(t *testing.T) {
			assert.Boolean(t, true, containsSegment("r", "r"))
		},
	)
	t.Run(
		"SingleLetterNoMatch",
		func(t *testing.T) {
			assert.Boolean(t, false, containsSegment("fooBar", "r"))
		},
	)
	t.Run(
		"MultiCharacterPrefixMatch",
		func(t *testing.T) {
			assert.Boolean(t, true, containsSegment("fooReference", "ref"))
		},
	)
	t.Run(
		"MultiCharacterExactMatch",
		func(t *testing.T) {
			assert.Boolean(t, true, containsSegment("fooRef", "ref"))
		},
	)
	t.Run(
		"MultiCharacterNoMatch",
		func(t *testing.T) {
			assert.Boolean(t, false, containsSegment("fooBar", "ref"))
		},
	)
	t.Run(
		"MultiCharacterTooShort",
		func(t *testing.T) {
			assert.Boolean(t, false, containsSegment("fooRe", "ref"))
		},
	)
}
