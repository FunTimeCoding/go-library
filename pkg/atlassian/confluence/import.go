package confluence

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/basic_client/response"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/page"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/page/page_file"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/page/page_post"
	"github.com/funtimecoding/go-library/pkg/notation"
	"github.com/funtimecoding/go-library/pkg/system"
)

func (c *Client) Import(
	space string,
	parent string,
	path string,
) *page.Page {
	f := page_file.Decode(system.ReadFile(path))
	var result *response.Page
	notation.DecodeStrict(
		c.basic.PostV2(
			"/pages",
			page_post.New(
				c.SpaceByName(space).Identifier,
				c.PageBySpaceAndName(space, parent).Identifier,
				f.Name,
				page.ToMarkup(f.Body),
			).Encode(),
		),
		&result,
		false,
	)

	return page.New(result, c.host)
}
