package cobra

import "github.com/spf13/cobra"

func New(name string) *cobra.Command {
	return &cobra.Command{Use: name}
}
