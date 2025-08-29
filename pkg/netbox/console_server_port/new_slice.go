package console_server_port

import "github.com/netbox-community/go-netbox/v4"

func NewSlice(v []netbox.ConsoleServerPort) []*Port {
	var result []*Port

	for _, e := range v {
		result = append(result, New(&e))
	}

	return result
}
