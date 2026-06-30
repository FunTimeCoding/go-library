package push

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/strings/upper"
	"testing"
)

func TestPush(t *testing.T) {
	assert.NotNil(t, New(upper.Alfa, 0, false, upper.Bravo))
}
