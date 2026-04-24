package gohab

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/habitica"
	"github.com/spf13/cobra"
)

func tasksCommand(c *habitica.Client) *cobra.Command {
	var taskType string
	command := &cobra.Command{
		Use:   "tasks",
		Short: "List tasks",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			fmt.Println(c.Tasks(taskType))
		},
	}
	command.Flags().StringVar(
		&taskType,
		"type",
		"",
		"Task type: habits, dailys, todos, rewards",
	)

	return command
}
