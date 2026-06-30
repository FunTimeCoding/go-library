package metric

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/strings/upper"
	"testing"
)

func TestMetric(t *testing.T) {
	assert.NotNil(t, New(upper.Alfa))
}
