package guilds

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/gw2/constant"
	"github.com/funtimecoding/go-library/pkg/notation"
	"github.com/funtimecoding/go-library/pkg/system"
)

func Parse(path string) map[string][]string {
	var result map[string][]string
	s := system.ReadFile(path, constant.GuildFile)

	if false {
		fmt.Printf("Parsing: %s\n", s)
	}

	notation.DecodeStrict(s, &result, true)

	return result
}
