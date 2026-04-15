package assert

import (
	"reflect"
	"testing"
)

func NotEmpty(
	t *testing.T,
	countable any,
) {
	t.Helper()
	v := reflect.ValueOf(countable)

	switch v.Kind() {
	case reflect.Slice, reflect.Array, reflect.Map:
		if v.Len() == 0 {
			t.Errorf("expected non-empty %T, got 0 items", countable)
		}
	default:
		t.Errorf("not countable: %T", countable)
	}
}
