package convert

import (
	"github.com/funtimecoding/go-library/pkg/habitica/response"
	"github.com/funtimecoding/go-library/pkg/tool/gohabd/generated/server"
)

func statsPointer(s response.Stats) *server.Stats {
	result := Stats(s)

	return &result
}
