package convert

import (
	"github.com/funtimecoding/go-library/pkg/habitica/response"
	"github.com/funtimecoding/go-library/pkg/tool/gohabiticad/generated/server"
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
		result.Checklist = new(ChecklistItems(t.Checklist))
	}

	if t.Streak > 0 {
		result.Streak = &t.Streak
	}

	if t.Value != 0 {
		result.Value = new(float32(t.Value))
	}

	return result
}
