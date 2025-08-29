package push

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/strings"
	"testing"
)

func TestPush(t *testing.T) {
	assert.NotNil(t, New(strings.Alfa, 0, false, strings.Bravo))
}
