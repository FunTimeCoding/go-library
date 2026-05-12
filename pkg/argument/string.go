package argument

func (i *Instance) String(
	name string,
	value string,
	usage string,
) {
	i.flags.String(name, value, usage)
}
