package split

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestSplit(t *testing.T) {
	assert.Strings(t, []string{"a", "b"}, Colon("a:b"))
	assert.Strings(t, []string{"a", "b"}, Comma("a,b"))
	assert.Strings(t, []string{"a", "b"}, Dash("a-b"))
	assert.Strings(t, []string{"a", "b"}, DoubleColon("a::b"))
	assert.Strings(t, []string{"a", "b"}, NewLine("a\nb"))
	assert.Strings(t, []string{"a", "b"}, Pipe("a|b"))
	assert.Strings(t, []string{"a", "b"}, Space("a b"))
	assert.Strings(t, []string{"a", "b"}, Underscore("a_b"))
}

func TestSplitEquals(t *testing.T) {
	a, b := Equals("a=b")
	assert.String(t, "a", a)
	assert.String(t, "b", b)

	c, d := Equals("c=d=e")
	assert.String(t, "c", c)
	assert.String(t, "d=e", d)

	f, g := Equals("f")
	assert.String(t, "f", f)
	assert.String(t, "", g)

	h, i := Equals("h=")
	assert.String(t, "h", h)
	assert.String(t, "", i)
}
