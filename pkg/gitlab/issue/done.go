package issue

func (i *Issue) Done() bool {
	return i.State == "closed"
}
