package node

import (
	"github.com/funtimecoding/go-library/pkg/strings/separator"
	"strings"
)

func (n *Node) FormatPath() string {
	var result []string

	for _, p := range n.GetPath() {
		result = append(result, p.Name)
	}

	return strings.Join(result, separator.Slash)
}
