package strings

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestDeleteDuplicates(t *testing.T) {
	assert.Any(
		t,
		[]string{Alfa, Bravo},
		DeleteDuplicates([]string{Alfa, Alfa, Bravo}),
	)
}
