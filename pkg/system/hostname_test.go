package system

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHostname(t *testing.T) {
	assert.True(t, len(Hostname()) > 0)
}
