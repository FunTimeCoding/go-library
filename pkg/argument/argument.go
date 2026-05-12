package argument

func (i *Instance) Argument(n int) string {
	return i.flags.Arg(n)
}
