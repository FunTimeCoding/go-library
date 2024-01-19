package system

const (
	NullPhysicalAddressString = "00:00:00:00:00:00"

	Linux  = "linux"
	Darwin = "darwin"

	AMD64 = "amd64"
	ARM64 = "arm64"

	DownloadsPath               = "Downloads"
	KubernetesConfigurationPath = ".kube/config"
	QuerySocketPath             = ".osquery/shell.em"
)

var NullPhysicalAddress = PhysicalAddress(NullPhysicalAddressString)
