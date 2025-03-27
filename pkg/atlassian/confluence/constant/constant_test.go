package constant

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestClient(t *testing.T) {
	assert.String(t, "no space", NoSpace)
}
