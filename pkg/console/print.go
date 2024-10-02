package console

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/console/format"
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/separator"
	"strings"
)

func Print(
	f face.Formattable,
	title string,
	indent int,
	s *format.Settings,
) {
	fmt.Printf(
		"%s%s: %s\n",
		strings.Repeat(separator.DoubleSpace, indent),
		title,
		f.Format(s),
	)
}
