package system

const (
	Tilde = "~"

	NullPhysicalAddressString = "00:00:00:00:00:00"

	Linux   = "linux"
	Darwin  = "darwin"
	Windows = "windows"

	AMD64 = "amd64"
	ARM64 = "arm64"

	Temporary                   = "tmp"
	DownloadsPath               = "Downloads"
	KubernetesConfigurationPath = ".kube/config"
	QuerySocketPath             = ".osquery/shell.em"
)

var NullPhysicalAddress = PhysicalAddress(NullPhysicalAddressString)
