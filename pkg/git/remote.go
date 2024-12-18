package git

import (
	"github.com/funtimecoding/go-library/pkg/git/remote"
	"github.com/funtimecoding/go-library/pkg/git/remote/provider_map"
)

func Remote(
	path string,
	m *provider_map.Map,
	name string,
) *remote.Remote {
	for _, element := range Remotes(path, m) {
		if element.Name == name {
			return element
		}
	}

	return nil
}
