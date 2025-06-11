package brave

import (
	"github.com/funtimecoding/go-library/pkg/brave/helper"
	"github.com/funtimecoding/go-library/pkg/brave/preference"
	"github.com/funtimecoding/go-library/pkg/brave/profile"
	"github.com/funtimecoding/go-library/pkg/system"
)

func Profiles() []*profile.Profile {
	var result []*profile.Profile

	for _, f := range system.ReadDirectory(helper.SettingsPath()) {
		if !IsProfile(f) {
			continue
		}

		p := profile.Parse(preference.Parse(f.Name()).Cookie.Accounts)[0]
		p.Profile = f.Name()
		result = append(result, p)
	}

	return result
}
