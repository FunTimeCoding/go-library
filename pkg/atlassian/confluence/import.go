package confluence

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/basic/response"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/constant"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/page"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/page/page_file"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/page/page_post"
	"github.com/funtimecoding/go-library/pkg/notation"
	"github.com/funtimecoding/go-library/pkg/system"
)

func (c *Client) Import(
	space string,
	parent string,
	base string,
	name string,
) (*page.Page, error) {
	f := page_file.Decode(system.ReadFile(base, name))
	s, e := c.SpaceByName(space)

	if e != nil {
		return nil, e
	}

	p, g := c.PageBySpaceAndName(space, parent)

	if g != nil {
		return nil, g
	}

	body, h := c.basic.PostV2Path(
		constant.Page,
		page_post.New(
			s.Identifier,
			p.Identifier,
			f.Name,
			page.ToMarkup(f.Body),
		).Encode(),
	)

	if h != nil {
		return nil, h
	}

	var result *response.Page
	notation.DecodeStrict(body, &result, false)

	return page.New(result, c.host), nil
}
