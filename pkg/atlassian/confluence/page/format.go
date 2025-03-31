package page

import (
	"github.com/funtimecoding/go-library/pkg/console/status"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
)

func (p *Page) Format(f *option.Format) string {
	return status.New(f).String(p.Name).Format()
}
