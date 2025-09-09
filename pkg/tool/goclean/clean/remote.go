package clean

import (
	"github.com/funtimecoding/go-library/pkg/git"
	"github.com/funtimecoding/go-library/pkg/git/constant"
	"github.com/funtimecoding/go-library/pkg/git/remote"
	"github.com/funtimecoding/go-library/pkg/git/remote/provider_map"
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/funtimecoding/go-library/pkg/tool/goclean/clean/option"
)

func originRemote(
	o *option.Clean,
	m *provider_map.Map,
) *remote.Remote {
	result := git.Remote(system.WorkingDirectory(), m, constant.OriginRemote)

	if result == nil {
		system.Exitf(
			1,
			"could not identify provider: %s\n",
			o.GitLabHost,
		)
		// make static analyzer happy
		result = &remote.Remote{}
	}

	return result
}
