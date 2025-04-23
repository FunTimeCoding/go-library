package validator

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestValidator(t *testing.T) {
	assert.True(t, New(nil) != nil)
}
