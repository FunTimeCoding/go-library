package metric

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/strings"
	"testing"
)

func TestMetric(t *testing.T) {
	assert.True(t, New(strings.Alfa) != nil)
}
