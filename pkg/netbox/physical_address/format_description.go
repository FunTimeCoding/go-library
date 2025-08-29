package physical_address

func (a *Address) formatDescription() string {
	return a.Raw.GetDescription()
}
