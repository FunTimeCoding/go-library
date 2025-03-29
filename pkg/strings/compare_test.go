package strings

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestCompare(t *testing.T) {
	assertCompare(
		t,
		[]string{},
		[]string{},
		[]string{},
		[]string{},
		[]string{},
	)
	// Add
	assertCompare(
		t,
		[]string{Alfa},
		[]string{},
		[]string{},
		[]string{},
		[]string{Alfa},
	)
	// Remove
	assertCompare(
		t,
		[]string{},
		[]string{Alfa},
		[]string{},
		[]string{Alfa},
		[]string{},
	)
	// Stay
	assertCompare(
		t,
		[]string{},
		[]string{},
		[]string{Alfa},
		[]string{Alfa},
		[]string{Alfa},
	)
}

func assertCompare(
	t *testing.T,
	expectAdd []string,
	expectRemove []string,
	expectStay []string,
	past []string,
	now []string,
) {
	t.Helper()
	add, remove, stay := Compare(past, now)
	assert.Any(t, expectAdd, add)
	assert.Any(t, expectRemove, remove)
	assert.Any(t, expectStay, stay)
}
