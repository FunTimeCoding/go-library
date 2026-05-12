package argument

func (i *Instance) Integer64(
	name string,
	value int64,
	usage string,
) {
	i.flags.Int64(name, value, usage)
}
