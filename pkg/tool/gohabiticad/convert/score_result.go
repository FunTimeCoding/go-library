package convert

import (
	"github.com/funtimecoding/go-library/pkg/habitica/score"
	"github.com/funtimecoding/go-library/pkg/tool/gohabiticad/generated/server"
)

func ScoreResult(r *score.Score) *server.ScoreResult {
	return &server.ScoreResult{
		Hp:    float32(r.HP),
		Mp:    float32(r.MP),
		Xp:    float32(r.XP),
		Gp:    float32(r.GP),
		Level: r.Level,
	}
}
