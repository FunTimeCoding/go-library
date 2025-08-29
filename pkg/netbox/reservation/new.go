package reservation

import "github.com/netbox-community/go-netbox/v4"

func New(d *netbox.RackReservation) *Reservation {
	return &Reservation{Identifier: d.GetId(), Name: d.GetDisplay(), Raw: d}
}
