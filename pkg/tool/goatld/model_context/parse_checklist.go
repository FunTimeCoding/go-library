package model_context

import (
	"github.com/funtimecoding/go-library/pkg/tool/goatld/convert"
	"strings"
)

func ParseChecklist(value string) []convert.ChecklistItem {
	var result []convert.ChecklistItem

	for _, line := range strings.Split(value, "\n") {
		line = strings.TrimSpace(line)

		if line == "" {
			continue
		}

		if strings.HasPrefix(line, "+ ") {
			result = append(
				result,
				convert.ChecklistItem{
					Index:   len(result) + 1,
					Text:    line[2:],
					Checked: true,
				},
			)
		} else if strings.HasPrefix(line, "- ") {
			result = append(
				result,
				convert.ChecklistItem{
					Index:   len(result) + 1,
					Text:    line[2:],
					Checked: false,
				},
			)
		}
	}

	return result
}
