package netbox

import (
	"fmt"
	"net"
)

func AddressMask(
	i net.IP,
	m net.IPMask,
) string {
	ones, _ := m.Size()

	return fmt.Sprintf("%s/%d", i.String(), ones)
}
