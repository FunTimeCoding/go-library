package port_forward_state

type PortForwardState struct {
	Identifier string
	Resource   string
	Namespace  string
	LocalPort  int
	TargetPort int
	Stop       chan struct{}
}
