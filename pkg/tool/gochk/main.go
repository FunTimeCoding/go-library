package gochk

import (
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/tool/gochk/check"
	"github.com/spf13/pflag"
)

func Main() {
	pflag.String(
		argument.Port,
		"",
		"Port, multiple values separated by comma",
	)
	argument.ParseBind()
	check.Check()
}
