package system

import "net"

func Lookup(address string) []string {
	result, e := net.LookupAddr(address)

	if e != nil {
		return []string{}
	}

	return result
}
