package contains

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/strings"
	"testing"
)

func TestOne(t *testing.T) {
	assert.True(t, One([]string{strings.Alfa}, strings.Alfa))
	assert.False(t, One([]string{strings.Alfa}, strings.Bravo))
}
