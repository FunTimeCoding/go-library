package argument

func (i *Instance) PrintUsage() {
	i.flags.Usage()
}
