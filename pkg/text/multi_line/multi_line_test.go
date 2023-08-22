package multi_line

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestMultiLine(t *testing.T) {
	l := New()
	l.Add("a")
	l.Add("b")
	assert.String(t, "a\nb", l.Format())
}
