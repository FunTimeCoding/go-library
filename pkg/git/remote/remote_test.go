package remote

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/strings"
	"testing"
)

func TestRemote(t *testing.T) {
	assert.Any(
		t,
		&Remote{
			Name:     "Alfa",
			Locator:  "Bravo",
			Provider: "Charlie",
		},
		New(strings.Alfa, strings.Bravo, strings.Charlie),
	)
}
