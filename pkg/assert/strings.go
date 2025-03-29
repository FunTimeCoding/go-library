package assert

import (
	"reflect"
	"testing"
)

func Strings(
	t *testing.T,
	expect []string,
	actual []string,
) {
	if reflect.DeepEqual(actual, expect) {
		return
	}

	t.Helper()
	t.Errorf("\nExpect: %+q\nActual: %+q", expect, actual)
}
