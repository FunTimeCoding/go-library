package argument

func (i *Instance) BooleanShort(
	name string,
	shorthand string,
	value bool,
	usage string,
) {
	i.flags.BoolP(name, shorthand, value, usage)
}
