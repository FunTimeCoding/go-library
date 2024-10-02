package format

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestCopy(t *testing.T) {
	alpha := New()
	alpha.Tags = []string{"a"}
	bravo := alpha.Copy()
	bravo.Tags = append(bravo.Tags, "b")
	assert.Strings(t, []string{"a"}, alpha.Tags)
	assert.Strings(t, []string{"a", "b"}, bravo.Tags)

	if alpha == bravo {
		t.Errorf(
			"Expected copy to be different: %+v %+v",
			alpha,
			bravo,
		)
	}
}
