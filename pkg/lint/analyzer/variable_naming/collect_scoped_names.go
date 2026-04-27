package variable_naming

import "go/types"

func collectScopedNames(o types.Object) map[string]bool {
	result := map[string]bool{}
	own := o.Parent()

	for s := own; s != nil; s = s.Parent() {
		for _, name := range s.Names() {
			if s == own && name == o.Name() {
				continue
			}

			result[name] = true
		}
	}

	return result
}
