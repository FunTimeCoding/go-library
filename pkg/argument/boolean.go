package argument

func (i *Instance) Boolean(
	name string,
	value bool,
	usage string,
) {
	i.flags.Bool(name, value, usage)
}
