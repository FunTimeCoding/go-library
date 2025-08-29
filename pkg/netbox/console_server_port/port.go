package console_server_port

import "github.com/netbox-community/go-netbox/v4"

type Port struct {
	Identifier int32
	Name       string

	Raw *netbox.ConsoleServerPort
}
