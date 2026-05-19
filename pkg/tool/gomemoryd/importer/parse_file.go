package importer

import (
	"os"
	"strings"
)

func ParseFile(path string) (*ParsedMemory, error) {
	raw, e := os.ReadFile(path)

	if e != nil {
		return nil, e
	}

	text := string(raw)
	m := &ParsedMemory{}

	if strings.HasPrefix(text, "---\n") {
		parts := strings.SplitN(text, "---\n", 3)

		if len(parts) >= 3 {
			parseFrontmatter(parts[1], m)
			m.Content = strings.TrimSpace(parts[2])
		} else {
			m.Content = strings.TrimSpace(text)
		}
	} else {
		m.Content = strings.TrimSpace(text)
	}

	if m.Description == "" {
		m.Description = m.Name
	}

	if m.Type == "" {
		m.Type = "feedback"
	}

	return m, nil
}
