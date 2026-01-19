package assert

import (
	"reflect"
	"testing"
)

func Nil(
	t *testing.T,
	actual any,
) {
	if actual == nil {
		return
	}

	t.Helper()
	v := reflect.ValueOf(actual)

	switch v.Kind() {
	case reflect.Ptr, reflect.Slice, reflect.Map, reflect.Chan, reflect.Func, reflect.Interface:
		if !v.IsNil() {
			t.Errorf("expected nil, got %T", actual)
		}
	default:
		t.Errorf("expected nil, got %T", actual)
	}
}
