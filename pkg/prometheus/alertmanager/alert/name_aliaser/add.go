package name_aliaser

func (a *Aliaser) Add(
	name string,
	alias string,
) *Aliaser {
	a.aliases[name] = alias

	return a
}
