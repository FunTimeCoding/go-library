package n8n

import (
	"github.com/funtimecoding/go-library/pkg/generative/n8n/response"
	"github.com/funtimecoding/go-library/pkg/notation"
)

func (c *Client) Workflows() []*response.Workflow {
	var result *response.Workflows
	notation.DecodeStrict(c.Get("/workflows"), &result, false)

	return result.Data
}
