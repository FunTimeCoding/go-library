package port_forward_state

func New(
	identifier string,
	resource string,
	namespace string,
	localPort int,
	targetPort int,
	stop chan struct{},
) *PortForwardState {
	return &PortForwardState{
		Identifier: identifier,
		Resource:   resource,
		Namespace:  namespace,
		LocalPort:  localPort,
		TargetPort: targetPort,
		Stop:       stop,
	}
}
