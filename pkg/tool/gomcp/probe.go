package gomcp

import (
	"context"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/mark3labs/mcp-go/client"
	"github.com/mark3labs/mcp-go/client/transport"
	"github.com/mark3labs/mcp-go/mcp"
)

func probe(
	url string,
	token string,
) {
	var opts []transport.StreamableHTTPCOption

	if token != "" {
		opts = append(
			opts,
			transport.WithHTTPHeaders(
				map[string]string{
					"Authorization": fmt.Sprintf("Bearer %s", token),
				},
			),
		)
	}

	c, e := client.NewStreamableHttpClient(url, opts...)
	errors.PanicOnError(e)

	ctx := context.Background()
	errors.PanicOnError(c.Start(ctx))

	_, e = c.Initialize(ctx, mcp.InitializeRequest{})
	errors.PanicOnError(e)

	fmt.Printf("Session ID: %s\n", c.GetTransport().GetSessionId())

	result, e := c.ListTools(ctx, mcp.ListToolsRequest{})
	errors.PanicOnError(e)

	for _, t := range result.Tools {
		fmt.Printf("%s: %s\n", t.Name, t.Description)
	}
}
