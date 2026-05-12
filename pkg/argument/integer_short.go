package argument

func (i *Instance) IntegerShort(
	name string,
	shorthand string,
	value int,
	usage string,
) {
	i.flags.IntP(name, shorthand, value, usage)
}
