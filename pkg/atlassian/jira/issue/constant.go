package issue

import "regexp"

const (
	NoStatus      = "no status"
	NoDescription = "no description"

	BlockedBy = "is blocked by"
)

var KeyMatch = regexp.MustCompile(`[A-Z]+-\d+`)
