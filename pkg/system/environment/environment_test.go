package environment

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/strings/upper"
	"testing"
)

func TestEnvironment(t *testing.T) {
	assert.String(
		t,
		"Alfa",
		Fallback("DOES_NOT_EXIST", upper.Alfa),
	)
	EnsureUnset("NEVER_EXIST")
}
