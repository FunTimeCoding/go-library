package face

import "github.com/funtimecoding/go-library/pkg/console/format"

type Formattable interface {
	Format(s *format.Settings) string
}
