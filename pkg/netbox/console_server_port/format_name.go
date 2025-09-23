package console_server_port

import (
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
)

func (p *Port) formatName(f *option.Format) string {
	if f.UseColor {
		return console.Cyan("%s", p.Name)
	}

	return p.Name
}
