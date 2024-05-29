package assert

import (
	"reflect"
	"testing"
)

func Integers(
	t *testing.T,
	expected []int,
	actual []int,
) {
	if reflect.DeepEqual(actual, expected) {
		return
	}

	t.Helper()
	t.Errorf(
		"\nExpected: %+q"+
			"\nActual:   %+q",
		expected,
		actual,
	)
}
