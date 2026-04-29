package raid_parser

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goraidparsed/generated/client"
	"github.com/funtimecoding/go-library/pkg/web"
	"time"
)

func (c *Client) Generate(
	files []string,
	date *time.Time,
) string {
	result, e := c.client.PostGenerate(
		c.context,
		client.PostGenerateJSONRequestBody{
			Files: files,
			Date:  date,
		},
	)
	errors.PanicOnError(e)

	return web.ReadString(result)
}
