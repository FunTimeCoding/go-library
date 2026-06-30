package strings

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/strings/upper"
	"testing"
)

func TestDeleteDuplicates(t *testing.T) {
	assert.Any(
		t,
		[]string{upper.Alfa, upper.Bravo},
		DeleteDuplicates([]string{upper.Alfa, upper.Alfa, upper.Bravo}),
	)
}
