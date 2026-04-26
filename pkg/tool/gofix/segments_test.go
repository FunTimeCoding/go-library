package gofix

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestSegments(t *testing.T) {
	t.Run(
		"CamelCase",
		func(t *testing.T) {
			assert.Strings(t, []string{"foo", "ref"}, segments("fooRef"))
		},
	)
	t.Run(
		"PascalCase",
		func(t *testing.T) {
			assert.Strings(
				t,
				[]string{"foo", "bar", "ref"},
				segments("FooBarRef"),
			)
		},
	)
	t.Run(
		"SnakeCase",
		func(t *testing.T) {
			assert.Strings(t, []string{"foo", "ref"}, segments("foo_ref"))
		},
	)
	t.Run(
		"SingleWord",
		func(t *testing.T) {
			assert.Strings(t, []string{"ref"}, segments("ref"))
		},
	)
	t.Run(
		"SingleUpperWord",
		func(t *testing.T) {
			assert.Strings(t, []string{"ref"}, segments("Ref"))
		},
	)
	t.Run(
		"SingleLetter",
		func(t *testing.T) {
			assert.Strings(t, []string{"r"}, segments("r"))
		},
	)
}

func TestSegmentSpans(t *testing.T) {
	t.Run(
		"CamelCase",
		func(t *testing.T) {
			spans := segmentSpans("fooRef")
			assert.Integer(t, 2, len(spans))
			assert.Integer(t, 0, spans[0].start)
			assert.Integer(t, 3, spans[0].end)
			assert.String(t, "foo", spans[0].lower)
			assert.Integer(t, 3, spans[1].start)
			assert.Integer(t, 6, spans[1].end)
			assert.String(t, "ref", spans[1].lower)
		},
	)
	t.Run(
		"SnakeCase",
		func(t *testing.T) {
			spans := segmentSpans("foo_ref")
			assert.Integer(t, 2, len(spans))
			assert.Integer(t, 0, spans[0].start)
			assert.Integer(t, 3, spans[0].end)
			assert.Integer(t, 4, spans[1].start)
			assert.Integer(t, 7, spans[1].end)
		},
	)
	t.Run(
		"SingleWord",
		func(t *testing.T) {
			spans := segmentSpans("ref")
			assert.Integer(t, 1, len(spans))
			assert.Integer(t, 0, spans[0].start)
			assert.Integer(t, 3, spans[0].end)
			assert.String(t, "ref", spans[0].lower)
		},
	)
}
