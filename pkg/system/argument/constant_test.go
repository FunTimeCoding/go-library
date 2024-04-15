package argument

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestConstant(t *testing.T) {
	assert.String(t, "branch", Branch)
	assert.String(t, "executable", Executable)
	assert.String(t, "header", Header)
	assert.String(t, "host", Host)
	assert.String(t, "locator", Locator)
	assert.String(t, "message", Message)
	assert.String(t, "name", Name)
	assert.String(t, "output", Output)
	assert.String(t, "owner", Owner)
	assert.String(t, "package", Package)
	assert.String(t, "password", Password)
	assert.String(t, "path", Path)
	assert.String(t, "project", Project)
	assert.String(t, "replace", Replace)
	assert.String(t, "repository", Repository)
	assert.String(t, "tag", Tag)
	assert.String(t, "template", Template)
	assert.String(t, "token", Token)
	assert.String(t, "user", User)
	assert.String(t, "verbose", Verbose)
	assert.String(t, "version", Version)
}
