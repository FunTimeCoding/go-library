package argument

func (i *Instance) StringVariable(
	p *string,
	name string,
	value string,
	usage string,
) {
	i.flags.StringVar(p, name, value, usage)
}
