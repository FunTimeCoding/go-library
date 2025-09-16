package environment

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/strings"
	"testing"
)

func TestEnvironment(t *testing.T) {
	assert.String(
		t,
		"Alfa",
		Default("DOES_NOT_EXIST", strings.Alfa),
	)

	EnsureUnset("NEVER_EXIST")
}
