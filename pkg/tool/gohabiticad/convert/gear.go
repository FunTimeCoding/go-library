package convert

import (
	"github.com/funtimecoding/go-library/pkg/habitica/response"
	"github.com/funtimecoding/go-library/pkg/tool/gohabiticad/generated/server"
)

func Gear(g response.Gear) server.GearResult {
	owned := make([]string, 0, len(g.Owned))

	for key, active := range g.Owned {
		if active {
			owned = append(owned, key)
		}
	}

	return server.GearResult{
		Equipped: g.Equipped,
		Owned:    owned,
	}
}
