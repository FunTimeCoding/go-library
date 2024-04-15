package debian

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestConstant(t *testing.T) {
	assert.String(t, "control", ControlFile)
	assert.String(t, "DEBIAN", PackageConfigurationDirectory)
	assert.String(t, ".deb", PackageExtension)
}
