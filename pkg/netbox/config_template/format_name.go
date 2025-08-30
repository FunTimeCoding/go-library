package config_template

import (
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
)

func (t *Template) formatName(f *option.Format) string {
	result := t.Name

	if f.UseColor {
		result = console.Cyan(result)
	}

	return result
}
