package network

func FindByName(
	i []*Interface,
	name string,
) *Interface {
	for _, element := range i {
		if element.Name == name {
			return element
		}
	}

	return nil
}
