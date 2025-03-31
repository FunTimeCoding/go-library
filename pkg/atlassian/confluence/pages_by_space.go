package confluence

import "fmt"

func (c *Client) PagesBySpace(identifier string) {
	fmt.Println(
		c.basic.GetV2(fmt.Sprintf("/spaces/%s/pages", identifier)),
	)
}
