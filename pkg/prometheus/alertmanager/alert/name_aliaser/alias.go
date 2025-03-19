package name_aliaser

func (a *Aliaser) Alias(name string) string {
	if alias, okay := a.aliases[name]; okay {
		return alias
	}

	return name
}
