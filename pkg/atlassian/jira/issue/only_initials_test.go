package issue

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestOnlyInitials(t *testing.T) {
	assert.Any(
		t,
		[]*Issue{{Initials: "ALFA"}},
		OnlyInitials(
			[]*Issue{{Initials: "ALFA"}, {Initials: "BRAVO"}},
			"ALFA",
		),
	)
}
