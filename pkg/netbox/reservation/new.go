package reservation

import "github.com/netbox-community/go-netbox/v4"

func New(v *netbox.RackReservation) *Reservation {
	return &Reservation{Identifier: v.GetId(), Name: v.GetDisplay(), Raw: v}
}
