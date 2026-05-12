package example

import (
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/generative/openai/site"
)

func Open() {
	a := argument.NewSimple("brave-open")
	a.ParseSimple()
	site.New().OpenChat(a.RequiredPositional(0, "CHAT"))
}
