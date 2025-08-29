package spy_writer

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestWriter(t *testing.T) {
	assert.NotNil(t, New())
}
