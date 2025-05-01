package main

import (
	"context"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/system/environment"
	"github.com/sashabaranov/go-openai"
)

func alternate() {
	c := openai.NewClient(environment.Get(TokenEnvironment, 1))
	resp, e := c.CreateChatCompletion(
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
	fmt.Println(resp.Choices[0].Message.Content)
}
