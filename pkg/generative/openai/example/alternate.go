package example

import (
	"context"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/generative/openai/constant"
	"github.com/funtimecoding/go-library/pkg/system/environment"
	"github.com/sashabaranov/go-openai"
)

func Alternate() {
	c := openai.NewClient(environment.Required(constant.TokenEnvironment))
	r, e := c.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT4o,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "Hello!",
				},
			},
		},
	)
	errors.PanicOnError(e)
	fmt.Println(r.Choices[0].Message.Content)
}
