package errors

import (
	"github.com/funtimecoding/go-library/assert"
	"testing"
)

func TestFormat(t *testing.T) {
	assert.String(t, `a: "b"`, Format("a", "b").Error())
}
