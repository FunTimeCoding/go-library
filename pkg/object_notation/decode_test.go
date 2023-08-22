package object_notation

import (
	assert2 "github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestDecode(t *testing.T) {
	var actual []int
	assert2.FatalOnError(t, Decode("[1]", &actual))
	assert2.Any(t, []int{1}, actual)
}

func TestDecodeStrict(t *testing.T) {
	var actual []int
	DecodeStrict("[1]", &actual)
	assert2.Any(t, []int{1}, actual)
}
