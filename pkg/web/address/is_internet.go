package address

import "net"

func IsInternet(s string) bool {
	return net.ParseIP(s) != nil
}
