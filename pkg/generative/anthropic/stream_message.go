package anthropic

import (
	"context"
	"github.com/anthropics/anthropic-sdk-go"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/generative/anthropic/message"
	"log"
)

func (c *Client) StreamMessage(
	ctx context.Context,
	messages []*message.Message,
	system string,
	model string,
) <-chan string {
	ch := make(chan string, 64)

	go func() {
		defer close(ch)

		stream := c.client.Messages.NewStreaming(
			ctx,
			anthropic.MessageNewParams{
				Model:     anthropic.Model(model),
				MaxTokens: 8192,
				System: []anthropic.TextBlockParam{
					{
						Text:         system,
						CacheControl: anthropic.NewCacheControlEphemeralParam(),
					},
				},
				Messages: message.ToParamSlice(messages),
			},
		)

		for stream.Next() {
			event := stream.Current()

			if event.Type == "content_block_delta" && event.Delta.Type == "text_delta" {
				select {
				case ch <- event.Delta.Text:
				case <-ctx.Done():
					return
				}
			}
		}

		if e := stream.Err(); e != nil {
			log.Printf("stream error: %v", e)
		}

		errors.LogClose(stream)
	}()

	return ch
}
