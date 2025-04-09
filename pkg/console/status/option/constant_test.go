package option

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConstant(t *testing.T) {
	assert.NotNil(t, Color)
	assert.NotNil(t, ExtendedColor)
}
