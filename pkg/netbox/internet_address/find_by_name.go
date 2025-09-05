package internet_address

func FindByName(
	v []*Address,
	name string,
) *Address {
	for _, e := range v {
		if e.Name == name {
			return e
		}
	}

	return nil
}
