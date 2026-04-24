package convert

import (
	"github.com/funtimecoding/go-library/pkg/tool/gohabd/habitica"
	"github.com/funtimecoding/go-library/pkg/tool/gohabd/server"
)

func ScoreResult(r habitica.ScoreResult) server.ScoreResult {
	return server.ScoreResult{
		Hp:    float32(r.HP),
		Mp:    float32(r.MP),
		Xp:    float32(r.XP),
		Gp:    float32(r.GP),
		Level: r.Level,
	}
}
