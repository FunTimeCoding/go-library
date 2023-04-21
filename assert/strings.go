package assert

import (
	"reflect"
	"testing"
)

func Strings(
	t *testing.T,
	expected []string,
	actual []string,
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
