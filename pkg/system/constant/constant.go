package constant

import "github.com/funtimecoding/go-library/pkg/network"

const (
	Tilde = "~"

	NullPhysicalAddressString = "00:00:00:00:00:00"

	Linux   = "linux"
	Darwin  = "darwin"
	Windows = "windows"

	AMD64 = "amd64"
	ARM64 = "arm64"

	LinuxAMD64  = "linux-amd64"
	DarwinARM64 = "darwin-arm64"
	DarwinAMD64 = "darwin-amd64"

	Command                     = "cmd"
	DownloadsPath               = "Downloads"
	KubernetesConfigurationPath = ".kube/config"
	QuerySocketPath             = ".osquery/shell.em"

	Transmission = "tcp"
)

// File hierarchy system
const (
	Binary    = "bin"
	Library   = "lib"
	Local     = "local"
	Resources = "usr"
	Source    = "src"
	Temporary = "tmp"
)

var (
	OperatingSystems    = []string{Linux, Darwin}
	Architectures       = []string{AMD64, ARM64}
	SystemArchitectures = []string{LinuxAMD64, DarwinARM64, DarwinAMD64}
)

var NullPhysicalAddress = network.PhysicalAddress(NullPhysicalAddressString)
