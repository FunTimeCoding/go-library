package opsgenie

import "github.com/funtimecoding/go-library/pkg/face"

func (c *Client) TagToTeam(f face.SliceAlias) {
	c.tagToTeam = f
}
