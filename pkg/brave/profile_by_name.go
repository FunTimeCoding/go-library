package brave

import "github.com/funtimecoding/go-library/pkg/brave/profile"

func ProfileByName(name string) *profile.Profile {
	for _, p := range Profiles() {
		if p.Profile == name {
			return p
		}
	}

	return nil
}
