package web

import "fmt"

func PortMap(
	from int,
	to int,
) string {
	ValidatePort(from)
	ValidatePort(to)

	return fmt.Sprintf("%d:%d", from, to)
}
