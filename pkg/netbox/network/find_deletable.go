package network

func FindDeletable(
	known []*Interface,
	expected []*Definition,
) []*Interface {
	var result []*Interface

	for _, i := range known {
		var found bool

		for _, inner := range expected {
			if i.Name == inner.Name {
				found = true
			}
		}

		if !found {
			result = append(result, i)
		}
	}

	return result
}
