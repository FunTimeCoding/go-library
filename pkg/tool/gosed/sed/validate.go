package sed

import "github.com/funtimecoding/go-library/pkg/tool/gosed/sed/option"

func validate(o *option.Sed) {
	if o.Host == "" {
		panic("host required")
	}

	if o.Token == "" {
		panic("token required")
	}

	if o.Owner == "" {
		panic("owner required")
	}

	if o.Repository == "" {
		panic("repository required")
	}

	if o.Branch == "" {
		panic("branch required")
	}

	if o.Path == "" {
		panic("path required")
	}

	if o.Message == "" {
		panic("message required")
	}

	if len(o.Replaces) == 0 {
		panic("at least one replace required")
	}
}
