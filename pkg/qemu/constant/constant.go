package constant

const (
	AMD64Command = "qemu-system-x86_64"
	ARM64Command = "qemu-system-aarch64"

	KVMAccelerator = "kvm" // Available Linux
	HVFAccelerator = "hvf" // Available on Darwin
)
