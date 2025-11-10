package gw2

import "github.com/MrGunflame/gw2api"

func New(token string) *Client {
	// https://wiki.guildwars2.com/wiki/API:Main
	// Alternate:
	//  https://github.com/yasvisu/gw2api
	//  https://github.com/42atomys/gw2api-go
	client := gw2api.New()

	if token != "" {
		client = client.WithAccessToken(token)
	}

	return &Client{client: client}
}
