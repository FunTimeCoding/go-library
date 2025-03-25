package color_assignment

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/strings"
	"testing"
)

func TestAssignment(t *testing.T) {
	assert.True(t, New(strings.Alfa, nil) != nil)
}
