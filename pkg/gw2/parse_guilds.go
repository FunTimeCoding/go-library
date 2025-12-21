package gw2

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/gw2/log_manager"
	"github.com/funtimecoding/go-library/pkg/notation"
)

func ParseGuilds(s string) []*log_manager.Guild {
	var guilds map[string]any
	notation.DecodeStrict(s, &guilds, false)
	var result []*log_manager.Guild

	for k, v := range guilds {
		if v == nil {
			errors.Warning("no data: %s", k)

			continue
		}

		var guild log_manager.Guild

		if e := notation.Decode(
			notation.Encode(v, false),
			&guild,
		); e != nil {
			errors.PanicOnError(e)
		}

		result = append(result, &guild)
	}

	return result
}
