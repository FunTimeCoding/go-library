package assert

import (
	"reflect"
	"testing"
)

func Integers(
	t *testing.T,
	expect []int,
	actual []int,
) {
	if reflect.DeepEqual(actual, expect) {
		return
	}

	t.Helper()
	t.Errorf("\nExpect: %+q\nActual: %+q", expect, actual)
}
