package node

import 	"github.com/funtimecoding/go-library/pkg/errors"

func DirectoryByNameStrict(
	n *Node,
	name string,
) *Node {
	result := DirectoryByName(n, name)

	if result == nil {
		errors.NotFound(name)

		return Stub()
	}

	return result
}
