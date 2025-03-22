package console

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/separator"
	"strings"
)

func Print(
	v face.Formattable,
	title string,
	indent int,
	f *option.Format,
) {
	fmt.Printf(
		"%s%s: %s\n",
		strings.Repeat(separator.DoubleSpace, indent),
		title,
		v.Format(f),
	)
}
