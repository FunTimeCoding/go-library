package gitlab

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestWithProjects(t *testing.T) {
	assert.NotNil(t, WithProjects([]int{}))
}
