package strings

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestConstant(t *testing.T) {
	assert.Strings(
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
