package strings

import (
	"github.com/funtimecoding/go-library/assert"
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
	// Remain
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
	expectedAdd []string,
	expectedRemove []string,
	expectedRemain []string,
	past []string,
	now []string,
) {
	t.Helper()
	add, remove, remain := Compare(past, now)
	assert.Any(t, expectedAdd, add)
	assert.Any(t, expectedRemove, remove)
	assert.Any(t, expectedRemain, remain)
}
