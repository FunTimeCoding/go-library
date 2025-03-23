package prometheus_detail

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestDetail(t *testing.T) {
	assert.True(t, New() != nil)
}
