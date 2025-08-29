package network

func (i *Interface) formatType() string {
	result := string(i.Type)

	if result == "" {
		result = NoType
	}

	return result
}
