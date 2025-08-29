package internet_address

func (a *Address) IsIdentical(d *Definition) bool {
	if !a.Address.Equal(d.Address) {
		return false
	}

	if a.Network.Mask.String() != d.Mask.String() {
		return false
	}

	return true
}
