package clean

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/git/remote/provider_map"
	"github.com/funtimecoding/go-library/pkg/tool/goclean/clean/option"
	web "github.com/funtimecoding/go-library/pkg/web/constant"
	"strings"
)

func Run(o *option.Clean) {
	o.GitLabHost = strings.TrimPrefix(o.GitLabHost, web.SecurePrefix)
	m := providerMap(o)
	r := originRemote(o, m)

	switch r.Provider {
	case provider_map.GitHubProvider:
		Hub(r)
	case provider_map.GitLabProvider:
		Lab(o, r)
	case provider_map.UnknownProvider:
		// TODO: Consider deleting tags except latest locally and pushing them to the server
		fmt.Println("Unknown provider, nothing to clean")
	}
}
