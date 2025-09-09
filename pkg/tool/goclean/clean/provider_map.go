package clean

import (
	"github.com/funtimecoding/go-library/pkg/git/remote/provider_map"
	"github.com/funtimecoding/go-library/pkg/tool/goclean/clean/option"
	"github.com/funtimecoding/go-library/pkg/web/locator"
)

func providerMap(o *option.Clean) *provider_map.Map {
	result := provider_map.New()
	result.Add(o.GitLabHost, provider_map.GitLabProvider)

	if locator.HasDot(o.GitLabHost) {
		result.Add(
			locator.RemoveSubdomain(o.GitLabHost),
			provider_map.GitLabProvider,
		)
	}

	return result
}
