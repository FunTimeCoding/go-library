package mark

import (
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/generative/model_context/example/mark/option"
)

func Main() {
	a := argument.NewSimple("mark-example")
	a.Boolean(argument.Local, false, "Local STDIO")
	a.ParseSimple()
	o := option.New()
	o.Local = a.GetBoolean(argument.Local)
	Run(o)
}
