package argument

import (
	"github.com/funtimecoding/go-library/pkg/identity"
	"github.com/spf13/pflag"
)

type Instance struct {
	identity *identity.Tool
	flags    *pflag.FlagSet
}
