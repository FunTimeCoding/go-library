package silence

func FilterPermanent(v []*Silence) []*Silence {
	var result []*Silence

	for _, s := range v {
		if !s.IsPermanent() {
			result = append(result, s)
		}
	}

	return result
}
