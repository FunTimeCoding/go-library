package ceil

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestInteger(t *testing.T) {
	assert.Integer(t, 1, Integer(0.9))
}
