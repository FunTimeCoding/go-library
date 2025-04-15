package issue

import "regexp"

const (
	NoStatus      = "no status"
	NoDescription = "no description"
)

var KeyMatch = regexp.MustCompile(`[A-Z]+-\d+`)
