package confluence

import "fmt"

func (c *Client) Delete(pageIdentifier string) {
	c.basic.DeleteV2(fmt.Sprintf("/pages/%s", pageIdentifier))
}
