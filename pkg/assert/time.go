package assert

import (
	"testing"
	"time"
)

func Time(
	t *testing.T,
	expect time.Time,
	actual time.Time,
) {
	t.Helper()
	expectRounded := expect.Round(time.Second)
	actualRounded := actual.Round(time.Second)

	if !actualRounded.Equal(expectRounded) {
		t.Errorf(
			"\nExpect: %#v\nActual: %#v",
			expectRounded,
			actualRounded,
		)
	}
}
