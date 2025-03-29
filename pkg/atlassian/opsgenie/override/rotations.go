package override

func (o *Override) rotations() []string {
	var result []string

	for _, element := range o.Rotations {
		result = append(result, element.Name)
	}

	return result
}
