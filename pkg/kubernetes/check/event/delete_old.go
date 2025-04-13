package event

import (
	"github.com/funtimecoding/go-library/pkg/kubernetes/client"
	"github.com/funtimecoding/go-library/pkg/kubernetes/types/native/event"
	"time"
)

func deleteOld(
	c *client.Client,
	e *event.Event,
	age time.Duration,
) bool {
	if e.Age() > age {
		c.DeleteEventSimple(e)

		return true
	}

	return false
}
