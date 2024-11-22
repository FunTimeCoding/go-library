package constant

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestConstant(t *testing.T) {
	assert.String(t, "-T", NoPTYArgument)
	assert.String(t, "-tt", ForcePTYArgument)
	assert.String(t, "-v", VerboseArgument)
}
