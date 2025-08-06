package discord

import "fmt"

func (c *Client) PrintStatus() {
	fmt.Printf("User: %+v\n", c.User().Username)

	for _, guild := range c.Guilds() {
		fmt.Printf("Guild: %+v\n", guild.ID)

		for _, channel := range c.Channels(guild.ID) {
			fmt.Printf("Channel: %s\n", channel.Name)
		}
	}
}
