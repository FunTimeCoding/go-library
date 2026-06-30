package issue

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/strings/capital"
	"testing"
)

func TestOnlyInitials(t *testing.T) {
	assert.Any(
		t,
		[]*Issue{{Initials: "ALFA"}},
		OnlyInitials(
			[]*Issue{
				{Initials: capital.Alfa},
				{Initials: capital.Bravo},
			},
			capital.Alfa,
		),
	)
}
