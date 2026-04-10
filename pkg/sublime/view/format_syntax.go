package view

import "github.com/funtimecoding/go-library/pkg/console/status/option"

func (v *View) formatSyntax(_ *option.Format) string {
	if v.Syntax == "" {
		return NoSyntax
	}

	return v.Syntax
}
