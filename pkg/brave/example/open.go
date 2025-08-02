package example

import (
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/generative/openai/site"
)

func Open() {
	argument.ParseBind()
	site.New().OpenChat(argument.RequiredPositional(0, "CHAT"))
}
