package clean

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/git"
	"github.com/funtimecoding/go-library/pkg/git/constant"
	"github.com/funtimecoding/go-library/pkg/git/remote/provider_map"
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/funtimecoding/go-library/pkg/tool/goclean/clean/option"
	"github.com/funtimecoding/go-library/pkg/web/locator"
)

func Run(o *option.Clean) {
	m := provider_map.New()
	m.Add(o.GitLabHost, provider_map.GitLabProvider)

	if locator.HasDot(o.GitLabHost) {
		m.Add(
			locator.RemoveSubdomain(o.GitLabHost),
			provider_map.GitLabProvider,
		)
	}

	origin := git.Remote(system.WorkingDirectory(), m, constant.OriginRemote)

	if origin == nil {
		system.Exitf(
			1,
			"could not identify provider: %s\n",
			o.GitLabHost,
		)

		return
	}

	switch origin.Provider {
	case provider_map.GitHubProvider:
		Hub(origin)
	case provider_map.GitLabProvider:
		Lab(o, origin)
	case provider_map.UnknownProvider:
		// TODO: Consider deleting tags except latest locally and pushing them to the server
		fmt.Println("Unknown provider, nothing to clean")
	}
}
