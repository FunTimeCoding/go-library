package strings

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestDeleteEmpty(t *testing.T) {
	assert.Any(
		t, []string{"Alfa", "Bravo"},
		DeleteEmpty([]string{"", Alfa, "", Bravo, ""}),
	)
}
