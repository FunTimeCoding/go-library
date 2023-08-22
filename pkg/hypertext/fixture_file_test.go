package hypertext

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestFixtureFile(t *testing.T) {
	_, e := fixtureFile("test.html")
	assert.FatalOnError(t, e)
}
