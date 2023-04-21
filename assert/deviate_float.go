package assert

import "testing"

func Deviate(
	t *testing.T,
	expected float64,
	actual float64,
	deviation float64,
) {
	amount := expected * deviation

	if amount < 0 {
		amount *= -1
	}

	if actual <= expected+amount && actual >= expected-amount {
		return
	}

	t.Helper()
	t.Errorf(
		"\nExpected: %g+-%g"+
			"\nActual:   %g",
		expected,
		amount,
		actual,
	)
}
