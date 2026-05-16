package tunnel_termination

import "github.com/netbox-community/go-netbox/v4"

type Termination struct {
	Identifier      int32
	TerminationType string
	TerminationIdentifier   int64
	Role            string
	Raw             *netbox.TunnelTermination
}
