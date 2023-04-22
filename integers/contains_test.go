package integers

import (
	"github.com/funtimecoding/go-library/assert"
	"testing"
)

func TestContains(t *testing.T) {
	assert.Boolean(
		t,
		true,
		Contains(0, []int{0}),
	)
	assert.Boolean(
		t,
		false,
		Contains(1, []int{0}),
	)
}
