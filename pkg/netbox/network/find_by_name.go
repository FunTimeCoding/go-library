package network

func FindByName(
	v []*Interface,
	name string,
) *Interface {
	for _, i := range v {
		if i.Name == name {
			return i
		}
	}

	return nil
}
