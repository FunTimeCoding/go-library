package network

func FindDeletable(
	known []*Interface,
	expected []*Definition,
) []*Interface {
	var result []*Interface

	for _, element := range known {
		var found bool

		for _, inner := range expected {
			if element.Name == inner.Name {
				found = true
			}
		}

		if !found {
			result = append(result, element)
		}
	}

	return result
}
