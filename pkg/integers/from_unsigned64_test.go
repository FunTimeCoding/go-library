package integers

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestFromUnsigned64(t *testing.T) {
	assert.Integer(t, 0, FromUnsigned64(0))
}
