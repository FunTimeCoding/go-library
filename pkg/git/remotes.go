package git

import (
	"github.com/funtimecoding/go-library/pkg/git/remote"
	"github.com/funtimecoding/go-library/pkg/git/remote/provider_map"
)

func Remotes(
	path string,
	m *provider_map.Map,
) []*remote.Remote {
	var result []*remote.Remote

	for _, element := range RemotesRaw(path) {
		c := element.Config()
		locator := c.URLs[0]
		result = append(
			result,
			remote.New(
				c.Name,
				locator,
				m.FindProvider(ParseLocator(locator).Host),
			),
		)
	}

	return result
}
