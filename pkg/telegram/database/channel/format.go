package channel

import "fmt"

func (c *Channel) Format() string {
	return fmt.Sprintf("%d %s\n", c.Identifier, c.Name)
}
