package main

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/system/environment"
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

func official() {
	c := openai.NewClient(
		option.WithAPIKey(environment.Get(TokenEnvironment, 1)),
	)
	chatCompletion, e := c.Chat.Completions.New(
		context.Background(),
		openai.ChatCompletionNewParams{
			Messages: []openai.ChatCompletionMessageParamUnion{
				openai.UserMessage("Say this is a test"),
			},
			Model: openai.ChatModelGPT4o,
		},
	)
	errors.PanicOnError(e)
	println(chatCompletion.Choices[0].Message.Content)
}
