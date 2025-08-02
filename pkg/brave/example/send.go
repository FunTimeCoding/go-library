package example

import (
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/generative/openai/site"
)

func Send() {
	argument.ParseBind()
	site.New().Send(argument.RequiredPositional(0, "TEXT"))
}
