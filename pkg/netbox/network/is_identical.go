package network

func (i *Interface) IsIdentical(d *Definition) bool {
	if string(i.Type) != d.Type {
		return false
	}

	if i.PhysicalAddress.String() != d.PhysicalAddress.String() {
		return false
	}

	return true
}
