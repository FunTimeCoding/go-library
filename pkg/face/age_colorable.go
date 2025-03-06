package face

import "time"

type AgeColorable interface {
	Age() time.Duration
	AgeColorFunction() SprintFunction
	SetAgeColorFunction(SprintFunction)
}
