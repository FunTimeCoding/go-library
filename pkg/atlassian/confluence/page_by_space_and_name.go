package confluence

import "fmt"

func (c *Client) PageBySpaceAndName(
	spaceIdentifier string,
	name string,
) {
	fmt.Println(
		c.basic.GetV2(
			fmt.Sprintf(
				"/pages?space-id=%s&title=%s",
				spaceIdentifier,
				name,
			),
		),
	)
}
