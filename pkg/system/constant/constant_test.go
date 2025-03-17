package constant

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"net"
	"testing"
)

func TestConstant(t *testing.T) {
	assert.String(t, "cmd", CommandPath)
	assert.String(t, "tmp", Temporary)
	assert.String(t, "lib", Library)
	assert.String(t, "usr", Resources)
	assert.String(t, "local", Local)
	assert.String(t, "bin", Binary)
	assert.String(t, "src", Source)
	assert.String(t, "Downloads", DownloadsPath)
	assert.String(t, ".kube/config", KubernetesConfigurationPath)
	assert.String(t, ".osquery/shell.em", QuerySocketPath)
	assert.String(t, "amd64", AMD64)
	assert.String(t, "arm64", ARM64)
	assert.String(t, "linux-amd64", LinuxAMD64)
	assert.String(t, "darwin-arm64", DarwinARM64)
	assert.String(t, "darwin-amd64", DarwinAMD64)
	assert.Strings(t, []string{"linux", "darwin"}, OperatingSystems)
	assert.Strings(
		t,
		[]string{"linux-amd64", "darwin-arm64", "darwin-amd64"},
		SystemArchitectures,
	)
	assert.Any(
		t,
		net.HardwareAddr{0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		NullPhysicalAddress,
	)
}
