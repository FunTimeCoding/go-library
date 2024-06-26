package web

import "fmt"

func Link(
	host string,
	port int,
	secure bool,
) string {
	scheme := Scheme(secure)

	if port != 0 {
		return fmt.Sprintf("%s://%s:%d", scheme, host, port)
	}

	return fmt.Sprintf("%s://%s", scheme, host)
}
