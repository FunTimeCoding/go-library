package internet_address

import "net"

func FilterSubnet(
	v []*Address,
	n *net.IPNet,
) []*Address {
	var result []*Address

	for _, e := range v {
		if n.Contains(e.Address) {
			result = append(result, e)
		}
	}

	return result
}
