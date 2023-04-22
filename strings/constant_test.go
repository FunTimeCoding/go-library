package strings

import (
	"github.com/funtimecoding/go-library/assert"
	"testing"
)

func TestConstant(t *testing.T) {
	assert.Any(
		t,
		[]string{
			"Alfa",
			"Bravo",
			"Charlie",
			"Delta",
			"Echo",
			"Foxtrot",
			"Golf",
			"Hotel",
			"India",
			"Juliett",
			"Kilo",
			"Lima",
			"Mike",
		},
		[]string{
			Alfa,
			Bravo,
			Charlie,
			Delta,
			Echo,
			Foxtrot,
			Golf,
			Hotel,
			India,
			Juliett,
			Kilo,
			Lima,
			Mike,
		},
	)
}
