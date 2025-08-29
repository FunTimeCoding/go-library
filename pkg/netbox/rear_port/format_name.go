package rear_port

import (
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
)

func (p *Port) formatName(f *option.Format) string {
	result := p.Name

	if f.UseColor {
		result = console.Cyan(result)
	}

	return result
}
