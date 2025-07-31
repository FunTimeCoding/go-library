package constant

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestConstant(t *testing.T) {
	assert.Count(t, 3, Severities)
	assert.Count(t, 4, Statuses)
}
