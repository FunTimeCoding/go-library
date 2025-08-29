package provider_map

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestProvider(t *testing.T) {
	assert.NotNil(t, New())
}
