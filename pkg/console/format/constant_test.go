package format

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConstant(t *testing.T) {
	assert.NotNil(t, Plain)
	assert.NotNil(t, Color)
	assert.NotNil(t, Raw)
	assert.NotNil(t, RawColor)
	assert.NotNil(t, Extended)
	assert.NotNil(t, ExtendedColor)
}
