package multi_line

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestMultiLine(t *testing.T) {
	l := New()
	assert.Boolean(t, true, l.Empty())
	l.Add("a")
	assert.Boolean(t, false, l.Empty())
	l.Add("b")
	assert.String(t, "a\nb", l.Format())
}
