package constant

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestConstant(t *testing.T) {
	assert.String(t, "gosensor", GoSensor)
	assert.String(t, "podman", PodManPrefix)
}
