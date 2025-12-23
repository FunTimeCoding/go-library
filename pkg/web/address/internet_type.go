package address

import "net"

func InternetType(s string) string {
	i := net.ParseIP(s)

	if i == nil {
		return NoneType
	}

	if i.To4() != nil {
		return FourType
	}

	return SixType
}
