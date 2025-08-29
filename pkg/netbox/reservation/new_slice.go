package reservation

import "github.com/netbox-community/go-netbox/v4"

func NewSlice(v []netbox.RackReservation) []*Reservation {
	var result []*Reservation

	for _, e := range v {
		result = append(result, New(&e))
	}

	return result
}
