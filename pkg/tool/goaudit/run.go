package goaudit

import (
	"github.com/funtimecoding/go-library/pkg/system/virtual_file_system"
	"github.com/funtimecoding/go-library/pkg/tool/goaudit/option"
	"github.com/funtimecoding/go-library/pkg/tool/goaudit/scan"
	"path/filepath"
)

func Run(o *option.Audit) {
	for _, root := range o.Roots {
		v := virtual_file_system.From(root)
		repo := filepath.Base(root)
		services := scan.Services(v, repo)
		identityWarnings := scan.IdentityWarnings(v, services)
		clients := scan.Clients(
			v,
			repo,
			scan.LoadConfiguration(".goaudit.yaml"),
		)

		if o.Table {
			runTable(services, identityWarnings, clients)

			continue
		}

		runHeadless(v, services, identityWarnings)
	}
}
