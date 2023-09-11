package assert

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/pmezard/go-difflib/difflib"
	"github.com/sanity-io/litter"
	"reflect"
	"testing"
)

func Any(
	t *testing.T,
	expected any,
	actual any,
) {
	t.Helper()

	if (expected == nil && actual != nil) ||
		(expected != nil && actual == nil) {
		t.Fatal(
			spew.Sdump(
				"\nNil comparison"+
					"\nExpected: %#v"+
					"\nActual:   %#v",
				expected,
				actual,
			),
		)
	} else if expected == nil && actual == nil {
		return
	}

	expectedLength := -1
	actualLength := -1

	switch reflect.TypeOf(expected).Kind() {
	case reflect.Slice:
		expectedLength = reflect.ValueOf(expected).Len()
	case reflect.Array:
		expectedLength = reflect.ValueOf(expected).Len()
	}

	switch reflect.TypeOf(actual).Kind() {
	case reflect.Slice:
		actualLength = reflect.ValueOf(actual).Len()
	case reflect.Array:
		actualLength = reflect.ValueOf(actual).Len()
	}

	if expectedLength > -1 || actualLength > -1 {
		var reason string

		if actualLength < expectedLength {
			reason = "Actual is shorter than expected"
		} else if actualLength > expectedLength {
			reason = "Actual is longer than expected"
		}

		if reason != "" {
			diff := difflib.UnifiedDiff{
				A:        difflib.SplitLines(litter.Sdump(expected)),
				B:        difflib.SplitLines(litter.Sdump(actual)),
				FromFile: fmt.Sprintf("Expected (%d)", expectedLength),
				ToFile:   fmt.Sprintf("Actual (%d)", actualLength),
				Context:  1,
			}
			text, _ := difflib.GetUnifiedDiffString(diff)
			t.Fatalf("%s\n%s", reason, text)
		}
	}

	if reflect.DeepEqual(actual, expected) {
		return
	}

	diff := difflib.UnifiedDiff{
		A:       difflib.SplitLines(spew.Sdump(expected)),
		B:       difflib.SplitLines(spew.Sdump(actual)),
		Context: 10,
	}
	text, _ := difflib.GetUnifiedDiffString(diff)
	t.Error(text)
}
