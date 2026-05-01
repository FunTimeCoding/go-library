package convert

import (
	"github.com/funtimecoding/go-library/pkg/habitica/response"
	"github.com/funtimecoding/go-library/pkg/tool/gohabd/generated/server"
)

func Task(t response.Task) server.Task {
	result := server.Task{
		Identifier: t.ID,
		Text:       t.Text,
		Type:       t.Type,
	}

	if t.Notes != "" {
		result.Notes = &t.Notes
	}

	if len(t.Tags) > 0 {
		result.Tags = &t.Tags
	}

	if t.Completed {
		result.Completed = &t.Completed
	}

	if len(t.Checklist) > 0 {
		items := ChecklistItems(t.Checklist)
		result.Checklist = &items
	}

	if t.Streak > 0 {
		result.Streak = &t.Streak
	}

	if t.Value != 0 {
		v := float32(t.Value)
		result.Value = &v
	}

	return result
}
