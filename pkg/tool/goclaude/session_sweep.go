package goclaude

import (
	"context"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/generated/client"
	"github.com/spf13/cobra"
)

func sessionSweep(c *client.ClientWithResponses) *cobra.Command {
	return &cobra.Command{
		Use:   "sweep",
		Short: "Copy session files to safe harbor",
		Args:  cobra.NoArgs,
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			response, e := c.PostSweepWithResponse(
				context.Background(),
			)
			errors.PanicOnError(e)
			r := response.JSON200
			fmt.Printf(
				"sweep: %d copied, %d updated, %d skipped\n",
				r.Copied,
				r.Updated,
				r.Skipped,
			)
		},
	}
}
