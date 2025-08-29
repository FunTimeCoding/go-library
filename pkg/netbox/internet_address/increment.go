package internet_address

import "net"

func increment(i net.IP) {
	for j := len(i) - 1; j >= 0; j-- {
		i[j]++

		if i[j] > 0 {
			break
		}
	}
}
