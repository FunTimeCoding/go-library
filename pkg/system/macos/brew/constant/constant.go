package constant

import "github.com/funtimecoding/go-library/pkg/console/status/option"

const (
	Brew      = "brew"
	Outdated  = "outdated"
	Info      = "info"
	Installed = "--installed"
	Notation2 = "--json=v2"
	Notation1 = "--json=v1"
)

var Format = option.Color.Copy()
