package assert

import (
	"github.com/funtimecoding/go-library/pkg/math/round"
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

	actual = round.Round(actual, decimals)

	// Pass if identical after rounding
	if actual == expected {
		return
	}

	t.Helper()
	var format string

	switch decimals {
	case 0:
		format = "\nExpected: %.0f\nActual:   %.0f"
	case 1:
		format = "\nExpected: %.1f\nActual:   %.1f"
	case 2:
		format = "\nExpected: %.2f\nActual:   %.2f"
	case 3:
		format = "\nExpected: %.3f\nActual:   %.3f"
	case 4:
		format = "\nExpected: %.4f\nActual:   %.4f"
	case 5:
		format = "\nExpected: %.5f\nActual:   %.5f"
	case 6:
		format = "\nExpected: %.6f\nActual:   %.6f"
	case 7:
		format = "\nExpected: %.7f\nActual:   %.7f"
	case 8:
		format = "\nExpected: %.8f\nActual:   %.8f"
	case 9:
		format = "\nExpected: %.9f\nActual:   %.9f"
	case 10:
		format = "\nExpected: %.10f\nActual:   %.10f"
	default:
		t.Error("Unexpected decimals:", decimals)

		return
	}

	t.Errorf(format, expected, actual)
}
