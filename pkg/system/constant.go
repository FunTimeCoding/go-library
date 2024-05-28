package system

const (
	Tilde = "~"

	NullPhysicalAddressString = "00:00:00:00:00:00"

	Linux   = "linux"
	Darwin  = "darwin"
	Windows = "windows"

	AMD64 = "amd64"
	ARM64 = "arm64"

	Command                     = "cmd"
	DownloadsPath               = "Downloads"
	KubernetesConfigurationPath = ".kube/config"
	QuerySocketPath             = ".osquery/shell.em"

	Transmission = "tcp"
)

// File hierarchy system
const (
	Temporary = "tmp"
	Library   = "lib"
	Resources = "usr"
	Local     = "local"
	Binary    = "bin"
)

var Architectures = []string{AMD64, ARM64}

var NullPhysicalAddress = PhysicalAddress(NullPhysicalAddressString)
