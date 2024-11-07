package web

func HostPortFromLocator(s string) string {
	return ParseLocator(s).Host
}
