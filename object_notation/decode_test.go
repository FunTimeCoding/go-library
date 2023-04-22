package object_notation

import (
	"github.com/funtimecoding/go-library/assert"
	"testing"
)

func TestDecode(t *testing.T) {
	var actual []int
	assert.FatalOnError(t, Decode("[1]", &actual))
	assert.Any(t, []int{1}, actual)
}

func TestDecodeStrict(t *testing.T) {
	var actual []int
	DecodeStrict("[1]", &actual)
	assert.Any(t, []int{1}, actual)
}
