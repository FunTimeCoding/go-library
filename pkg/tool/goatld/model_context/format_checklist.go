package model_context

import (
	"github.com/funtimecoding/go-library/pkg/tool/goatld/convert"
	"strings"
)

func formatChecklist(items []convert.ChecklistItem) string {
	var lines []string

	for _, item := range items {
		if item.Checked {
			lines = append(lines, "+ "+item.Text)
		} else {
			lines = append(lines, "- "+item.Text)
		}
	}

	return strings.Join(lines, "\n") + "\n"
}
