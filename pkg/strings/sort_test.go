package strings

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestSort(t *testing.T) {
	preSorted := []string{Alfa, Bravo, Charlie}
	Sort(preSorted, true)
	assert.Strings(
		t,
		[]string{Alfa, Bravo, Charlie},
		preSorted,
	)

	ascending := []string{Bravo, Alfa, Charlie}
	Sort(ascending, true)
	assert.Strings(
		t,
		[]string{Alfa, Bravo, Charlie},
		ascending,
	)

	descending := []string{Bravo, Alfa, Charlie}
	Sort(descending, false)
	assert.Strings(
		t,
		[]string{Charlie, Bravo, Alfa},
		descending,
	)
}
