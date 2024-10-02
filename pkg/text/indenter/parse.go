package indenter

import "strings"

type Node struct {
	Text     string
	Children []*Node
}

func Parse(s string) *Node {
	root := &Node{}
	inputLines := strings.Split(s, "\n")
	parseLines(inputLines, root, 0)

	return root
}

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

func leadingSpaces(line string) int {
	for i, element := range line {
		if element != ' ' {
			return i
		}
	}

	return len(line)
}
