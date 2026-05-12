package argument

func (i *Instance) BooleanVariable(
	p *bool,
	name string,
	value bool,
	usage string,
) {
	i.flags.BoolVar(p, name, value, usage)
}
