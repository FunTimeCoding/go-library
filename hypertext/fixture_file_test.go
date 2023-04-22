package hypertext

import (
	"github.com/funtimecoding/go-library/assert"
	"testing"
)

func TestFixtureFile(t *testing.T) {
	_, e := fixtureFile("test.html")
	assert.FatalOnError(t, e)
}
