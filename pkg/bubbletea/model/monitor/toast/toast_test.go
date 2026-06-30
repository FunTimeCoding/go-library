package toast

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/strings/upper"
	"testing"
)

func TestToast(t *testing.T) {
	assert.NotNil(t, New(0, upper.Alfa))
}
