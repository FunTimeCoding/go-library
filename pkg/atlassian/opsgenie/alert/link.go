package alert

import "fmt"

func link(
	identifier string,
	webHost string,
) string {
	return fmt.Sprintf(
		"https://%s/alert/detail/%s/details",
		webHost,
		identifier,
	)
}
