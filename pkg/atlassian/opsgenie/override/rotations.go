package override

func (o *Override) rotations() []string {
	var result []string

	for _, e := range o.Rotations {
		result = append(result, e.Name)
	}

	return result
}
