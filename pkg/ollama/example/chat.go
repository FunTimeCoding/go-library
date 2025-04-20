package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/ollama"
	"github.com/funtimecoding/go-library/pkg/ollama/chat_request"
	"github.com/funtimecoding/go-library/pkg/ollama/constant"
	"github.com/ollama/ollama/api"
)

func Chat() {
	o := ollama.NewEnvironment()
	r := chat_request.New()
	r.Messages = append(
		r.Messages,
		api.Message{
			Role:    constant.SystemRole,
			Content: "You are a ChatOps bot. Respond in the users message length and writing style.",
		},
	)
	r.Messages = append(
		r.Messages,
		api.Message{
			Role:    constant.UserRole,
			Content: "My server is down.",
		},
	)
	fmt.Printf("Messages: %+v\n", r.Messages)
	e := o.Chat(r)
	fmt.Printf("Response 1: %+v\n", e)
	r2 := chat_request.New()
	r2.Messages = append(r2.Messages, r.Messages...)
	r2.Messages = append(r2.Messages, e.Message)
	r2.Messages = append(
		r2.Messages,
		api.Message{
			Role:    constant.UserRole,
			Content: "What else can you suggest?",
		},
	)
	e2 := o.Chat(r2)
	fmt.Printf("Response 2: %+v\n", e2)
}
