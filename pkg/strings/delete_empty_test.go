package strings

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/strings/upper"
	"testing"
)

func TestDeleteEmpty(t *testing.T) {
	assert.Any(
		t,
		[]string{"Alfa", "Bravo"},
		DeleteEmpty([]string{"", upper.Alfa, "", upper.Bravo, ""}),
	)
}
