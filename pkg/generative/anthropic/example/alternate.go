package example

import (
	"context"
	"errors"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/generative/anthropic/constant"
	"github.com/funtimecoding/go-library/pkg/system/environment"
	"github.com/liushuangls/go-anthropic"
)

func Alternate() {
	// https://github.com/liushuangls/go-anthropic
	c := anthropic.NewClient(environment.Get(constant.TokenEnvironment))
	r, e := c.CreateMessages(
		context.Background(),
		anthropic.MessagesRequest{
			Model: anthropic.ModelClaude3Haiku20240307,
			Messages: []anthropic.Message{
				anthropic.NewUserTextMessage("What is your name?"),
			},
			MaxTokens: 1000,
		},
	)

	if e != nil {
		var n *anthropic.APIError

		if errors.As(e, &n) {
			fmt.Printf(
				"Messages error, type: %s, message: %s",
				n.Type,
				n.Message,
			)
		} else {
			fmt.Printf("Messages error: %v\n", e)
		}

		return
	}

	fmt.Println(r.Content[0].Text)
}
