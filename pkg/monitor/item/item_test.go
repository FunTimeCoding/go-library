package item

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/strings"
	"testing"
)

func TestItem(t *testing.T) {
	assert.String(t, "Alfa", New(strings.Alfa).Identifier)
}
