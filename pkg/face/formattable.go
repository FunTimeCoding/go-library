package face

import (
	"github.com/funtimecoding/go-library/pkg/console/description"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
)

type Formattable interface {
	Format(f *option.Format) string
	Meta() *description.Description
}
