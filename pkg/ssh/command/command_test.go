package command

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/strings/upper"
	"testing"
)

func TestCommand(t *testing.T) {
	assert.NotNil(t, New(upper.Alfa))
}
