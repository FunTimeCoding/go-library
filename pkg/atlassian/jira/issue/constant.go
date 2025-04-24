package issue

import "regexp"

const (
	NoStatus      = "no status"
	NoDescription = "no description"

	BlockedBy = "is blocked by"

	NilValue = "nil value"

	UnknownField = "unknown field"
	UnknownValue = "unknown value"
)

var KeyMatch = regexp.MustCompile(`[A-Z]+-\d+`)
