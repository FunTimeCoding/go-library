package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/reservation"
)

func (c *Client) MustReservations() []*reservation.Reservation {
	result, e := c.Reservations()
	errors.PanicOnError(e)

	return result
}
