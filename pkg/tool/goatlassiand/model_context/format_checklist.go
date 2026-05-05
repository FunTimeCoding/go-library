package model_context

import (
	"github.com/funtimecoding/go-library/pkg/strings/join"
	"github.com/funtimecoding/go-library/pkg/tool/goatlassiand/convert"
	"strings"
)

func formatChecklist(items []convert.ChecklistItem) string {
	var lines []string

	for _, item := range items {
		if item.Checked {
			lines = append(lines, join.Empty("+ ", item.Text))
		} else {
			lines = append(lines, join.Empty("- ", item.Text))
		}
	}

	return join.Empty(strings.Join(lines, "\n"), "\n")
}
