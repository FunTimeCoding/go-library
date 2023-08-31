package system

const (
	NullPhysicalAddressString = "00:00:00:00:00:00"

	Linux  = "linux"
	Darwin = "darwin"

	AMD64 = "amd64"
	ARM64 = "arm64"
)

var NullPhysicalAddress = PhysicalAddress(NullPhysicalAddressString)
