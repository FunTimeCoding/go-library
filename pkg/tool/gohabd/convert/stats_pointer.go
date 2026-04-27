package convert

import (
	"github.com/funtimecoding/go-library/pkg/tool/gohabd/habitica/response"
	"github.com/funtimecoding/go-library/pkg/tool/gohabd/server"
)

func statsPointer(s response.Stats) *server.Stats {
	result := Stats(s)

	return &result
}
