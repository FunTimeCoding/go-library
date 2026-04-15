package variable_naming

import "go/types"

// Rename represents a variable that should be renamed.
type Rename struct {
	Object  types.Object
	NewName string
}
