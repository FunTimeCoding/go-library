package opsgenie

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/strings"
	"testing"
)

func TestOpsgenie(t *testing.T) {
	assert.True(t, New(strings.Alfa, strings.Bravo, strings.Charlie) != nil)
}
