package field_changer

func (a *Changer) AddName(
	from string,
	to string,
) *Changer {
	a.name[from] = to

	return a
}
