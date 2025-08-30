package system_number

import "github.com/netbox-community/go-netbox/v4"

func NewSlice(v []netbox.ASN) []*Number {
	var result []*Number

	for _, e := range v {
		result = append(result, New(&e))
	}

	return result
}
