package physical_address

import (
	"log"
	"slices"
)

func validateObjectType(objectType string) {
	if !slices.Contains(ObjectTypes, objectType) {
		log.Panicf("unexpected object type: %s", objectType)
	}
}
