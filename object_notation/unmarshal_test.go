package object_notation

import (
	"github.com/funtimecoding/go-library/assert"
	"testing"
)

func TestUnmarshal(t *testing.T) {
	var actual []int
	Unmarshal("[1]", &actual)
	assert.Any(t, []int{1}, actual)
}
