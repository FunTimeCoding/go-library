package mock_client

import "github.com/funtimecoding/go-library/pkg/habitica/response"

func New() *Client {
	return &Client{
		stats: response.Stats{
			HP:    50,
			MP:    30,
			XP:    100,
			GP:    10,
			Level: 5,
			Class: "warrior",
		},
		gear: response.Gear{
			Equipped: map[string]string{},
			Owned:    map[string]bool{},
		},
	}
}
