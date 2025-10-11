package n8n

import (
	"github.com/funtimecoding/go-library/pkg/generative/n8n/constant"
	"github.com/funtimecoding/go-library/pkg/generative/n8n/response"
	"github.com/funtimecoding/go-library/pkg/notation"
)

func (c *Client) Workflows() []*response.Workflow {
	var r *response.Workflows
	notation.DecodeStrict(c.Get(constant.Workflows), &r, false)

	return r.Data
}
