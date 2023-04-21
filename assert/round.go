package assert

import (
	"github.com/funtimecoding/go-library/math"
	"testing"
)

func Round(
	t *testing.T,
	expected float64,
	actual float64,
	decimals int,
) {
	// Pass if identical before rounding
	if actual == expected {
		return
	}

	actual = math.Round(actual, decimals)

	// Pass if identical after rounding
	if actual == expected {
		return
	}

	t.Helper()
	var format string

	if decimals == 0 {
		format = "\nExpected: %.0f\nActual:   %.0f"
	} else if decimals == 1 {
		format = "\nExpected: %.1f\nActual:   %.1f"
	} else if decimals == 2 {
		format = "\nExpected: %.2f\nActual:   %.2f"
	} else if decimals == 3 {
		format = "\nExpected: %.3f\nActual:   %.3f"
	} else if decimals == 4 {
		format = "\nExpected: %.4f\nActual:   %.4f"
	} else if decimals == 5 {
		format = "\nExpected: %.5f\nActual:   %.5f"
	} else if decimals == 6 {
		format = "\nExpected: %.6f\nActual:   %.6f"
	} else if decimals == 7 {
		format = "\nExpected: %.7f\nActual:   %.7f"
	} else if decimals == 8 {
		format = "\nExpected: %.8f\nActual:   %.8f"
	} else {
		t.Error("Unexpected decimals:", decimals)

		return
	}

	t.Errorf(format, expected, actual)
}
