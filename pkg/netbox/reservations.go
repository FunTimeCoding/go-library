package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/constant"
	"github.com/funtimecoding/go-library/pkg/netbox/reservation"
)

func (c *Client) Reservations() []*reservation.Reservation {
	result, _, e := c.client.DcimAPI.DcimRackReservationsList(
		c.context,
	).Limit(constant.PageLimit).Execute()
	errors.PanicOnError(e)

	return reservation.NewSlice(result.Results)
}
