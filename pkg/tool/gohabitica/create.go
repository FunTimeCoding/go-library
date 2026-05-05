package gohabitica

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/tool/gohabiticad/client"
	"github.com/spf13/cobra"
)

func create(c *client.Client) *cobra.Command {
	var taskType string
	var text string
	var notes string
	command := &cobra.Command{
		Use:   "create",
		Short: "Create a task",
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			fmt.Println(c.CreateTask(taskType, text, notes))
		},
	}
	command.Flags().StringVar(
		&taskType,
		"type",
		"",
		"Task type: habit, daily, todo, reward",
	)
	command.Flags().StringVar(&text, "text", "", "Task title")
	command.Flags().StringVar(&notes, "notes", "", "Task notes")
	command.MarkFlagsRequiredTogether("type", "text")

	return command
}
