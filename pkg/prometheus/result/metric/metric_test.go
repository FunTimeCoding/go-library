package metric

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/strings"
	"testing"
)

func TestMetric(t *testing.T) {
	assert.NotNil(t, New(strings.Alfa))
}
