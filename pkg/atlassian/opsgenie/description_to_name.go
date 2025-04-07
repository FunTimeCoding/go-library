package opsgenie

import "github.com/funtimecoding/go-library/pkg/face"

func (c *Client) DescriptionToName(f face.StringAlias) {
	c.descriptionToName = f
}
