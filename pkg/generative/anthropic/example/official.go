package example

import (
	"context"
	"fmt"
	"github.com/anthropics/anthropic-sdk-go"
	"github.com/anthropics/anthropic-sdk-go/option"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/generative/anthropic/constant"
	"github.com/funtimecoding/go-library/pkg/system/environment"
)

func Official() {
	// https://github.com/anthropics/anthropic-sdk-go
	c := anthropic.NewClient(
		option.WithAPIKey(environment.Get(constant.TokenEnvironment)),
	)
	r, e := c.Messages.New(
		context.Background(),
		anthropic.MessageNewParams{
			MaxTokens: 1024,
			Messages: []anthropic.MessageParam{
				{
					Role: anthropic.MessageParamRoleUser,
					Content: []anthropic.ContentBlockParamUnion{
						{
							OfText: &anthropic.TextBlockParam{
								Text: "Explain a cat in 10 words.",
							},
						},
					},
				},
			},
			Model: anthropic.ModelClaude3_7SonnetLatest,
		},
	)
	errors.PanicOnError(e)
	fmt.Printf("%+v\n", r.Content)
}
