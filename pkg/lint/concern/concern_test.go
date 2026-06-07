package concern

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/strings"
	"testing"
)

func TestNewLine(t *testing.T) {
	c := NewLine(
		strings.Alfa,
		strings.Bravo,
		strings.Charlie,
		1,
		strings.Charlie,
		false,
	)
	assert.NotNil(t, c)
	assert.String(t, Line, c.Type)
	assert.Integer(t, 1, c.Line)
}

func TestNewFile(t *testing.T) {
	c := NewFile(
		strings.Alfa,
		strings.Bravo,
		strings.Charlie,
		false,
	)
	assert.NotNil(t, c)
	assert.String(t, File, c.Type)
	assert.Integer(t, 0, c.Line)
}

func TestNewPackage(t *testing.T) {
	c := NewPackage(
		strings.Alfa,
		strings.Bravo,
		strings.Charlie,
	)
	assert.NotNil(t, c)
	assert.String(t, Package, c.Type)
	assert.Boolean(t, false, c.Fixed)
}

func TestNewDelegatesToNewLine(t *testing.T) {
	c := New(
		strings.Alfa,
		strings.Bravo,
		strings.Charlie,
		1,
		strings.Charlie,
		false,
	)
	assert.String(t, Line, c.Type)
}
