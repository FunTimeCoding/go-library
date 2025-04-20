package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/ollama"
	"github.com/funtimecoding/go-library/pkg/ollama/chat_request"
)

func Chat() {
	o := ollama.NewEnvironment()
	r := chat_request.New()
	r.System(
		"You are a ChatOps bot. Respond in the users message length and writing style.",
	)
	r.User("My server is down.")
	fmt.Printf("Messages: %+v\n", r.Get().Messages)
	e := o.Chat(r.Get())
	fmt.Printf("Response 1: %+v\n", e)
	r2 := chat_request.New()
	r2.Import(r)
	r2.Message(e.Message)
	r2.User("What else can you suggest?")
	fmt.Printf("Response 2: %+v\n", o.Chat(r2.Get()))
}
