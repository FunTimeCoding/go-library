package argument

func (i *Instance) Integer(
	name string,
	value int,
	usage string,
) {
	i.flags.Int(name, value, usage)
}
