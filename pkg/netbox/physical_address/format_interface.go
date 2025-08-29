package physical_address

func (a *Address) formatInterface() string {
	if a.Interface == nil {
		return ""
	}

	return a.Interface.GetDisplay()
}
