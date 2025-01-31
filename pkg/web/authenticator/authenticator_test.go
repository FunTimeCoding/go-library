package authenticator

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestAuthenticator(t *testing.T) {
	assert.True(t, New() != nil)
}
