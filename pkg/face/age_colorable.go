package face

import "time"

type AgeColorable interface {
	Age() time.Duration
	AgeColor() SprintFunction
	SetAgeColor(SprintFunction)
}
