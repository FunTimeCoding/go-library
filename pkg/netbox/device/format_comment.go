package device

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
	library "github.com/funtimecoding/go-library/pkg/constant"
	"github.com/funtimecoding/go-library/pkg/netbox/constant"
	"github.com/funtimecoding/go-library/pkg/strings"
	"github.com/funtimecoding/go-library/pkg/text"
)

func (d *Device) formatComment(f *option.Format) string {
	if d.Comment == constant.NoComment {
		return ""
	}

	o := text.OptimizeWhitespace(d.Comment, library.CompactWhitespace)

	if o == "" {
		return ""
	}

	var result string

	if strings.CountRune(o, '\n') > 0 {
		result = fmt.Sprintf(
			"\n%s",
			strings.Indent(o, 4, false),
		)
	} else {
		result = fmt.Sprintf(" %s", o)
	}

	if f.UseColor {
		result = console.Magenta("%s", result)
	}

	return result
}
