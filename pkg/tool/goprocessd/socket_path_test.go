package goprocessd

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"strings"
	"testing"
)

func TestSocketPathDeterministic(t *testing.T) {
	a := socketPath("/tmp/project/Procfile")
	b := socketPath("/tmp/project/Procfile")
	assert.String(t, a, b)
}

func TestSocketPathDifferentDirectories(t *testing.T) {
	a := socketPath("/tmp/alpha/Procfile")
	b := socketPath("/tmp/bravo/Procfile")
	assert.True(t, a != b)
}

func TestSocketPathEndsWith(t *testing.T) {
	result := socketPath("/tmp/project/Procfile")
	assert.True(t, strings.HasSuffix(result, ".sock"))
}
