package web

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/raid"
)

func fightMap(f raid.Fight) string {
	if f.MapName != "" {
		return f.MapName
	}

	return fmt.Sprintf("%d", f.MapID)
}
