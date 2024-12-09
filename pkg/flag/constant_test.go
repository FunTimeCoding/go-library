package flag

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestConstant(t *testing.T) {
	assert.String(t, "filter", FilterFlag)
	assert.String(t, "investigate", InvestigateFlag)
}
