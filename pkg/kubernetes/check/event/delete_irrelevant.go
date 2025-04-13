package event

import (
	"github.com/funtimecoding/go-library/pkg/kubernetes/client"
	"github.com/funtimecoding/go-library/pkg/kubernetes/constant"
	"github.com/funtimecoding/go-library/pkg/kubernetes/types/native/event"
	"slices"
)

func deleteIrrelevant(
	c *client.Client,
	e *event.Event,
) bool {
	if slices.Contains(constant.IrrelevantEventReason, e.Reason) {
		c.DeleteEventSimple(e)

		return true
	}

	return false
}
