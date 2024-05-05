package cobra

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func Execute(c *cobra.Command) {
	if e := c.Execute(); e != nil {
		fmt.Println(e)

		os.Exit(1)
	}
}
