package notation

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestMarshallIndent(t *testing.T) {
	assert.Any(t, `"a"`, MarshallIndent("a"))
}
