package issue

func Children(
	all []*Issue,
	parent string,
) []*Issue {
	var result []*Issue

	for _, k := range all {
		if p := k.Raw.Fields.Parent; p != nil {
			if p.Key == parent {
				result = append(result, k)
			}
		}
	}

	return result
}
