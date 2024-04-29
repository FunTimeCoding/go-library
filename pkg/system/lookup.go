package system

import "net"

func Lookup(address string) []string {
	result, e := net.LookupAddr(CleanAddress(address))

	if e != nil {
		return []string{}
	}

	return result
}
