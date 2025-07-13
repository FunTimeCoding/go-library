package constant

const (
	SocketEnvironment = "SSH_AUTH_SOCK"

	Command                  = "ssh"
	LocalPortForwardArgument = "-L"
	NoRemoteCommandArgument  = "-N"
	BackgroundArgument       = "-f"
	NoPTYArgument            = "-T"
	ForcePTYArgument         = "-tt"
	VerboseArgument          = "-v"
)
