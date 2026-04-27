package convert

import (
	"github.com/funtimecoding/go-library/pkg/tool/gohabd/habitica/response"
	"github.com/funtimecoding/go-library/pkg/tool/gohabd/server"
)

func Stats(s response.Stats) server.Stats {
	return server.Stats{
		Hp:    float32(s.HP),
		Mp:    float32(s.MP),
		Xp:    float32(s.XP),
		Gp:    float32(s.GP),
		Level: s.Level,
		Class: s.Class,
	}
}
