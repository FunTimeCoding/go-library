package option

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestCommit(t *testing.T) {
	assert.NotNil(t, New())
}
