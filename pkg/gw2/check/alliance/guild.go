package alliance

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/gw2"
	"github.com/funtimecoding/go-library/pkg/system"
)

func Guild(path string) {
	for _, guild := range gw2.ParseGuilds(system.ReadFile(path)) {
		fmt.Printf("Guild: %s\n", guild.Name)
	}
}
