package netbox

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSerial(t *testing.T) {
	fixture := "0123456"
	assert.Equal(t, "123456", fixture[len(fixture)-6:])
	assert.Equal(t, "23456", fixture[len(fixture)-5:])
	assert.Equal(t, "3456", fixture[len(fixture)-4:])
}
