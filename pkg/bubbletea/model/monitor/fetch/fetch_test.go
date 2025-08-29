package fetch

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestFetch(t *testing.T) {
	assert.NotNil(t, Command())
}
