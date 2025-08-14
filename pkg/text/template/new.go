package template

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"text/template"
)

func New(
	name string,
	t string,
) *template.Template {
	result, e := template.New(name).Parse(t)
	errors.PanicOnError(e)

	return result
}
