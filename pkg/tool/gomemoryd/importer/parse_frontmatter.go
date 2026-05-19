package importer

import "strings"

func parseFrontmatter(
	block string,
	m *ParsedMemory,
) {
	for _, line := range strings.Split(block, "\n") {
		line = strings.TrimSpace(line)

		if k, v, okay := strings.Cut(line, ":"); okay {
			k = strings.TrimSpace(k)
			v = strings.TrimSpace(v)

			switch k {
			case "name":
				m.Name = v
			case "description":
				m.Description = v
			case "type":
				m.Type = v
			}
		}
	}
}
