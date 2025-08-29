package reservation

import "github.com/netbox-community/go-netbox/v4"

type Reservation struct {
	Identifier int32
	Name       string

	Raw *netbox.RackReservation
}
