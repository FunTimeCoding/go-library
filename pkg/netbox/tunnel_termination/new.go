package tunnel_termination

import "github.com/netbox-community/go-netbox/v4"

func New(v *netbox.TunnelTermination) *Termination {
	return &Termination{
		Identifier:            v.GetId(),
		TerminationType:       v.GetTerminationType(),
		TerminationIdentifier: v.GetTerminationId(),
		Role:                  string(v.Role.GetValue()),
		Raw:                   v,
	}
}
