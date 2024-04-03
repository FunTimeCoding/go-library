package notation

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/strings"
	"testing"
)

func TestDecodeAny(t *testing.T) {
	var a any

	DecodeAny(true, &a)
	assert.Any(t, a, true)

	DecodeAny(strings.Alfa, &a)
	assert.Any(t, a, "Alfa")
}
