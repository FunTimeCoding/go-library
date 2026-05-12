package argument

func (i *Instance) Positionals() []string {
	return i.flags.Args()
}
