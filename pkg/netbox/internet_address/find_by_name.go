package internet_address

func FindByName(
	i []*Address,
	name string,
) *Address {
	for _, element := range i {
		if element.Name == name {
			return element
		}
	}

	return nil
}
