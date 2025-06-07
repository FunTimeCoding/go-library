package gmail

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestClient(t *testing.T) {
	assert.True(t, New("") != nil)
}
