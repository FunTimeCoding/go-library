package argument

func (i *Instance) IntegerVariable(
	p *int,
	name string,
	value int,
	usage string,
) {
	i.flags.IntVar(p, name, value, usage)
}
