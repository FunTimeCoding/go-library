package system

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"net"
	"testing"
)

func TestConstant(t *testing.T) {
	assert.String(t, "cmd", Command)
	assert.String(t, "tmp", Temporary)
	assert.String(t, "lib", Library)
	assert.String(t, "usr", Resources)
	assert.String(t, "local", Local)
	assert.String(t, "bin", Binary)
	assert.String(t, "Downloads", DownloadsPath)
	assert.String(t, ".kube/config", KubernetesConfigurationPath)
	assert.String(t, ".osquery/shell.em", QuerySocketPath)
	assert.String(t, "amd64", AMD64)
	assert.String(t, "arm64", ARM64)
	assert.Any(
		t,
		net.HardwareAddr{0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		NullPhysicalAddress,
	)
}
