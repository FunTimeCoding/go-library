package identity

import "strings"

func (t *Tool) RenderInstructions(conditions map[string]bool) string {
	var parts []string
	parts = append(parts, t.instructions)

	for _, p := range t.paragraphs {
		if conditions[p.Key] {
			parts = append(parts, p.Text)
		}
	}

	return strings.Join(parts, " ")
}
