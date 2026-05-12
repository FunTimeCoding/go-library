package argument

func (i *Instance) ArgumentCount() int {
	return i.flags.NArg()
}
