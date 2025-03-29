package assert

import "testing"

func Deviate(
	t *testing.T,
	expect float64,
	actual float64,
	deviation float64,
) {
	amount := expect * deviation

	if amount < 0 {
		amount *= -1
	}

	if actual <= expect+amount && actual >= expect-amount {
		return
	}

	t.Helper()
	t.Errorf("\nExpect: %g+-%g\nActual: %g", expect, amount, actual)
}
