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

	CommandPath                 = "cmd"
	ConfigurationPath           = ".config"
	DownloadsPath               = "Downloads"
	FixturePath                 = "fixture"
	IdeaPath                    = ".idea"
	KubernetesConfigurationPath = ".kube/config"
	NotationPath                = "notation"
	HypertextPath               = "hypertext"
	MarkdownPath                = "markdown"
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

// macOS specific paths
const (
	MacOSLibrary = "Library"
	MacOsLogs    = "Logs"
)

var (
	OperatingSystems    = []string{Linux, Darwin}
	Architectures       = []string{AMD64, ARM64}
	SystemArchitectures = []string{LinuxAMD64, DarwinARM64, DarwinAMD64}
	NullPhysicalAddress = network.PhysicalAddress(NullPhysicalAddressString)
)
