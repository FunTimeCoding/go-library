package assert

import (
	"testing"
	"time"
)

func Time(
	t *testing.T,
	expected time.Time,
	actual time.Time,
) {
	t.Helper()
	expectedRounded := expected.Round(time.Second)
	actualRounded := actual.Round(time.Second)

	if !actualRounded.Equal(expectedRounded) {
		t.Errorf(
			"\nExpected: %#v"+
				"\nActual:   %#v",
			expectedRounded,
			actualRounded,
		)
	}
}
