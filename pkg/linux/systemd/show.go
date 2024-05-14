package systemd

import "github.com/funtimecoding/go-library/pkg/linux/systemd/show"

func (c *Client) Show(name string) *show.Result {
	return show.Parse(show.OutputToMap(c.ShowOutput(name)))
}
