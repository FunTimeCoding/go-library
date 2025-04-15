package indenter

import "strings"

func Parse(s string) *Node {
	root := &Node{}
	inputLines := strings.Split(s, "\n")
	parseLines(inputLines, root, 0)

	return root
}
