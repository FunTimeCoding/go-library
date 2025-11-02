package brave

import (
	"github.com/funtimecoding/go-library/pkg/brave/helper"
	"github.com/funtimecoding/go-library/pkg/brave/preference"
	"github.com/funtimecoding/go-library/pkg/brave/profile"
	"github.com/funtimecoding/go-library/pkg/system"
)

func Profiles() []*profile.Profile {
	var result []*profile.Profile
	p := helper.SettingsPath()

	for _, f := range system.ReadDirectory(p) {
		if !profile.IsProfile(f) {
			continue
		}

		if r := preference.Parse(f.Name()); r.Cookie.Accounts != "" {
			result = append(
				result,
				profile.Parse(p, f.Name(), r.Cookie.Accounts),
			)
		}
	}

	return result
}
