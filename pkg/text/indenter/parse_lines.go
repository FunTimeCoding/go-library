package indenter

import "strings"

func parseLines(
	lines []string,
	parent *Node,
	currentIndent int,
) {
	var node *Node

	for len(lines) > 0 {
		line := lines[0]
		indent := leadingSpaces(line)
		t := strings.TrimSpace(line)

		if t == "" {
			lines = lines[1:]

			continue
		}

		if indent == currentIndent {
			node = &Node{Text: t, Children: []*Node{}}
			parent.Children = append(parent.Children, node)
			lines = lines[1:]
		} else if indent > currentIndent {
			parseLines(lines, node, indent)

			return
		} else {
			return
		}
	}
}
