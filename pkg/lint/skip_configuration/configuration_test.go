package skip_configuration

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestConfiguration(t *testing.T) {
	assert.True(t, New("", false) != nil)
}
