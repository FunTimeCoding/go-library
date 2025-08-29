package metric

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestMetric(t *testing.T) {
	assert.NotNil(t, New(0, false))
}
