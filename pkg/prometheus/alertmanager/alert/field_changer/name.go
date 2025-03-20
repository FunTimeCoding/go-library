package field_changer

func (a *Changer) Name(n string) string {
	if alias, okay := a.name[n]; okay {
		return alias
	}

	return n
}
