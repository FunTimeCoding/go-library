package mock_client

import "github.com/funtimecoding/go-library/pkg/habitica/response"

type Client struct {
	tasks []response.Task
	tags  []response.Tag
	stats response.Stats
	gear  response.Gear
}
