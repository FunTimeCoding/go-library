package system

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddress(t *testing.T) {
	assert.NotNil(t, Address())
}
