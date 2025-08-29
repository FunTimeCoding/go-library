package internet_address

import "net"

func FilterSubnet(
	a []*Address,
	n *net.IPNet,
) []*Address {
	var result []*Address

	for _, element := range a {
		if n.Contains(element.Address) {
			result = append(result, element)
		}
	}

	return result
}
