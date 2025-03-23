package option

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/strings"
	"testing"
)

func TestOption(t *testing.T) {
	assert.True(
		t,
		New(strings.Alfa, strings.Bravo, []string{}, nil) != nil,
	)
}
