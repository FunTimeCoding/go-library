package goanalyze

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestChooseFix(t *testing.T) {
	t.Run(
		"SingleSegmentReturnsFirst",
		func(t *testing.T) {
			assert.String(
				t,
				"r",
				chooseFix("ref", []string{"r", "reference"}),
			)
		},
	)
	t.Run(
		"MultiSegmentPrefersMultiCharacter",
		func(t *testing.T) {
			assert.String(
				t,
				"reference",
				chooseFix("fooRef", []string{"r", "reference"}),
			)
		},
	)
	t.Run(
		"MultiSegmentFallsBackToFirst",
		func(t *testing.T) {
			assert.String(t, "t", chooseFix("fooTx", []string{"t"}))
		},
	)
	t.Run(
		"SingleOptionReturnsIt",
		func(t *testing.T) {
			assert.String(
				t,
				"arguments",
				chooseFix("args", []string{"arguments"}),
			)
		},
	)
}
