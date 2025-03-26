package pair

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/strings"
	"testing"
)

func TestPair(t *testing.T) {
	assert.True(t, New(strings.Alfa, strings.Bravo) != nil)
}
