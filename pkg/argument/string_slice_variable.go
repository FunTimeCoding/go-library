package argument

func (i *Instance) StringSliceVariable(
	p *[]string,
	name string,
	value []string,
	usage string,
) {
	i.flags.StringSliceVar(p, name, value, usage)
}
