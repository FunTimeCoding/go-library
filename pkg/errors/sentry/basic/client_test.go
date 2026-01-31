package basic

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestClient(t *testing.T) {
	assert.NotNil(t, New("", ""))
}
