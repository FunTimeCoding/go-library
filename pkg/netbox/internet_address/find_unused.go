package internet_address

import "net"

func FindUnused(
	v []*Address,
	n *net.IPNet,
) net.IP {
	used := make(map[string]bool)

	for _, e := range v {
		used[e.Address.String()] = true
	}

	first := n.IP.Mask(n.Mask)
	increment(first)

	for i := first; n.Contains(i); increment(i) {
		if !used[i.String()] {
			return i
		}
	}

	return nil
}
